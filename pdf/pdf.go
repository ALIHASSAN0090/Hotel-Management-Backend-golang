package pdf

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/joho/godotenv"
	gomail "gopkg.in/mail.v2"
)

func GenerateAndSendPDF(htmlFile, pdfFile string) error {

	clientCompanyAddress := request.CompanyAddress
	clientCompanyEmail := request.CompanyEmail

	clientEmail := request.CustomerPersonalEmail
	customergovNumber := request.CustomergovIDNum
	CustomerFullLegalName := request.CustomerFullLegalName

	price := request.Decided_price
	totalPrice := request.TotalPrice

	date := getDate()

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

	err = GeneratePDF(htmlFile, pdfFile, *clientCompanyAddress, *clientCompanyEmail, *clientEmail, *customergovNumber, *CustomerFullLegalName, *price, *totalPrice, date, strings.Join(servicenames, ", "))
	if err != nil {
		return fmt.Errorf("failed to generate PDF: %v", err)
	}

	fmt.Println("PDF generated successfully!")

	err = SendEmailNotification(pdfFile, smtpHost, smtpPort, senderUserName, senderPassword, *clientEmail)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	fmt.Println("Email sent successfully")
	return nil
}

func getDate() string {
	currentTime := time.Now()

	return currentTime.Format("02-01-2006")
}

func GeneratePDF(htmlFile, outputPath string, clientCompanyAddress, clientCompanyEmail, clientEmail, customerNumber, CustomerFullLegalName string, price, totalPrice float64, date string, serviceName string) error {
	// Read the HTML content
	htmlContent, err := os.ReadFile(htmlFile)
	if err != nil {
		return fmt.Errorf("failed to read HTML file: %v", err)
	}

	// Optionally, modify the HTML content with the provided values
	// For example, you could replace placeholders in the HTML with these values
	htmlContentStr := string(htmlContent)

	htmlContentStr = strings.ReplaceAll(htmlContentStr, "{{ClientCompanyAddress}}", clientCompanyAddress)
	htmlContentStr = strings.ReplaceAll(htmlContentStr, "{{ClientCompanyEmail}}", clientCompanyEmail)
	htmlContentStr = strings.ReplaceAll(htmlContentStr, "{{ClientEmail}}", clientEmail)
	htmlContentStr = strings.ReplaceAll(htmlContentStr, "{{CustomerNumber}}", customerNumber)
	htmlContentStr = strings.ReplaceAll(htmlContentStr, "{{CustomerFullLegalName}}", CustomerFullLegalName)
	htmlContentStr = strings.ReplaceAll(htmlContentStr, "{{Price}}", fmt.Sprintf("%.2f", price))
	htmlContentStr = strings.ReplaceAll(htmlContentStr, "{{TotalPrice}}", fmt.Sprintf("%.2f", totalPrice))
	htmlContentStr = strings.ReplaceAll(htmlContentStr, "{{Date}}", date)
	htmlContentStr = strings.ReplaceAll(htmlContentStr, "{{ServiceName}}", serviceName)

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
	m.Attach(pdfFile) // Attach the PDF file

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
