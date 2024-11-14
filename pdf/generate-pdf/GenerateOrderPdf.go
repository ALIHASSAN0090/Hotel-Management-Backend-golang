package pdf

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/joho/godotenv"
)

func GenerateAndSendOrderPDF(htmlFile, pdfFile string, customerName string, customerEmail string, resDate, resTime string, numberOfPersons int64, totalPrice float64, allfoodNames string) error {

	err := godotenv.Load("./.env")
	if err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	smtpHost := os.Getenv("MAIL_HOST")
	smtpPort := os.Getenv("MAIL_PORT")
	senderUserName := os.Getenv("MAIL_USERNAME")
	senderPassword := os.Getenv("MAIL_PASSWORD")

	if smtpHost == "" || smtpPort == "" || senderUserName == "" || senderPassword == "" {
		return fmt.Errorf("SMTP configuration and recipient email must be set in the environment variables")
	}

	err = GenerateOrderPDF(htmlFile, pdfFile, customerName, customerEmail, resDate, resTime, numberOfPersons, totalPrice, allfoodNames)
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

func GenerateOrderPDF(htmlFile, outputPath string, customerName string, customerEmail string, resDate, resTime string, numberOfPersons int64, totalPrice float64, allfoods string) error {

	htmlContent, err := os.ReadFile(htmlFile)
	if err != nil {
		return fmt.Errorf("failed to read HTML file: %v", err)
	}

	htmlContentStr := string(htmlContent)

	htmlContentStr = strings.ReplaceAll(htmlContentStr, "{{CustomerFullLegalName}}", customerName)
	htmlContentStr = strings.ReplaceAll(htmlContentStr, "{{ClientEmail}}", customerEmail)

	// htmlContentStr = strings.ReplaceAll(htmlContentStr, "{{Reservation Date}}", resDate)
	// htmlContentStr = strings.ReplaceAll(htmlContentStr, "{{date}}", resTime)
	htmlContentStr = strings.ReplaceAll(htmlContentStr, "{{Price}}", fmt.Sprintf("%.2f", totalPrice))
	htmlContentStr = strings.ReplaceAll(htmlContentStr, "{{Date}}", GetCurrentDate())

	foodsList := strings.ReplaceAll(allfoods, ",", ",\n")
	htmlContentStr = strings.ReplaceAll(htmlContentStr, "{{Food Details}}", foodsList)
	// htmlContentStr = strings.ReplaceAll(htmlContentStr, "{{Number of persons}}", fmt.Sprintf("%d", numberOfPersons))

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
