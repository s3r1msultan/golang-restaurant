package controllers

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendMessage(to, subject, message string) error {
	from := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	auth := smtp.PlainAuth("", from, password, smtpHost)
	msg := []byte(fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\nContent-Type: text/html; charset=\"UTF-8\"\n\n%s",
		from, to, subject, message))
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
}
