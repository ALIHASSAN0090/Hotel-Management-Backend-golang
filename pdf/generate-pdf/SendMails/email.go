package pdf

import (
	"crypto/tls"
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendEmailNotification(pdfFile, smtpHost, smtpPort, senderUserName, senderPassword, recieverEmail string) error {

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
