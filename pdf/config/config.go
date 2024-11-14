package pdf

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadSMTPConfig() (string, string, string, string, error) {
	err := godotenv.Load("./.env")
	if err != nil {
		return "", "", "", "", fmt.Errorf("error loading .env file: %v", err)
	}

	smtpHost := os.Getenv("MAIL_HOST")
	smtpPort := os.Getenv("MAIL_PORT")
	senderUserName := os.Getenv("MAIL_USERNAME")
	senderPassword := os.Getenv("MAIL_PASSWORD")

	if smtpHost == "" || smtpPort == "" || senderUserName == "" || senderPassword == "" {
		return "", "", "", "", fmt.Errorf("SMTP configuration must be set in the environment variables")
	}

	return smtpHost, smtpPort, senderUserName, senderPassword, nil
}