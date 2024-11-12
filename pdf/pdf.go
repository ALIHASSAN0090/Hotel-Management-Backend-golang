package pdf

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopkg.in/gomail.v2"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/joho/godotenv"
)

func GenerateAndSendPDF(htmlFile, pdfFile string, customerName string, customerEmail string, resDate, resTime string, numberOfPersons int64, totalPrice float64) error {

	err := godotenv.Load("./.env")
	if err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	smtpHost := os.Getenv("MAIL_HOST")
	smtpPort := os.Getenv("MAIL_PORT")
	senderUserName := os.Getenv("MAIL_USERNAME")
	senderPassword := os.Getenv("MAIL_PASSWORD")
	// recieverEmail1 := os.Getenv("RECIEVER_EMAIL")

	if smtpHost == "" || smtpPort == "" || senderUserName == "" || senderPassword == "" {
		return fmt.Errorf("SMTP configuration and recipient email must be set in the environment variables")
	}

	err = GeneratePDF(htmlFile, pdfFile, customerName, customerEmail, resDate, resTime, numberOfPersons, totalPrice)
	if err != nil {
		return fmt.Errorf("failed to generate PDF: %v", err)
	}

	fmt.Println("PDF generated successfully!")

	err = SendEmailNotification(pdfFile, smtpHost, smtpPort, senderUserName, senderPassword, customerEmail)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	fmt.Println("Email sent successfully")
	return nil
}

func GeneratePDF(htmlFile, outputPath string, customerName string, customerEmail string, resDate, resTime string, numberOfPersons int64, totalPrice float64) error {
	// Read the HTML content
	htmlContent, err := os.ReadFile(htmlFile)
	if err != nil {
		return fmt.Errorf("failed to read HTML file: %v", err)
	}

	// Optionally, modify the HTML content with the provided values
	// For example, you could replace placeholders in the HTML with these values
	htmlContentStr := string(htmlContent)

	htmlContentStr = strings.ReplaceAll(htmlContentStr, "{{CustomerFullLegalName}}", customerName)
	htmlContentStr = strings.ReplaceAll(htmlContentStr, "{{ClientEmail}}", customerEmail)

	htmlContentStr = strings.ReplaceAll(htmlContentStr, "{{Reservation Date}}", resDate)
	htmlContentStr = strings.ReplaceAll(htmlContentStr, "{{Reservation Time}}", resTime)
	htmlContentStr = strings.ReplaceAll(htmlContentStr, "{{Price}}", fmt.Sprintf("%.2f", totalPrice))

	// htmlContentStr = strings.ReplaceAll(htmlContentStr, "{{Date}}", date)
	htmlContentStr = strings.ReplaceAll(htmlContentStr, "{{Number of persons}}", fmt.Sprintf("%d", numberOfPersons))

	// Set up Chrome options
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
	)

	// Create context
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	var buf []byte

	err = chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate("data:text/html," + htmlContentStr),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			buf, _, err = page.PrintToPDF().
				WithPrintBackground(true).
				WithMarginTop(0.0).
				WithMarginBottom(0.0).
				WithMarginLeft(0.0).
				WithMarginRight(0.0).
				WithPaperWidth(8.27).
				WithPaperHeight(11.69).
				Do(ctx)
			return err
		}),
	})
	if err != nil {
		return fmt.Errorf("failed to render PDF: %v", err)
	}

	// Write the PDF file
	err = os.WriteFile(outputPath, buf, 0644)
	if err != nil {
		return fmt.Errorf("failed to save PDF: %v", err)
	}

	return nil
}

func SendEmailNotification(pdfFile, smtpHost, smtpPort, senderUserName, senderPassword, recieverEmail string) error {
	// Check if the PDF file exists
	if _, err := os.Stat(pdfFile); os.IsNotExist(err) {
		return fmt.Errorf("PDF file not found: %v", err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", senderUserName)
	m.SetHeader("To", recieverEmail)
	m.SetHeader("Subject", "Your Invoice")
	m.SetBody("text/plain", "Please find the attached PDF document for your review.")
	m.Attach(pdfFile)

	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		return fmt.Errorf("invalid SMTP port: %v", err)
	}

	d := gomail.NewDialer(smtpHost, port, senderUserName, senderPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
