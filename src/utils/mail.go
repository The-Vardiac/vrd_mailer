package utils

import (
	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

type Mail struct {
	From       string
	To         string
	CcAddress  string
	CcName     string
	Subject    string
	Body       string
	Attachment string
}

func (mail Mail) SendMail() error {
	smtp_server := os.Getenv("SMTP_SERVER")
	smtp_port := os.Getenv("SMTP_PORT")
	smtp_username := os.Getenv("SMTP_USERNAME")
	smtp_password := os.Getenv("SMTP_PASSWORD")
	smtp_sender_email := os.Getenv("SMTP_SENDER_EMAIL")

	smtp_port_int, _ := strconv.Atoi(smtp_port)
	mail.From = smtp_sender_email

	newMessage := gomail.NewMessage()
	newMessage.SetHeader("From", mail.From)
	newMessage.SetHeader("To", mail.To)
	if mail.CcAddress != "" && mail.CcName != "" {
		newMessage.SetAddressHeader("Cc", mail.CcAddress, mail.CcName)
	}
	newMessage.SetHeader("Subject", mail.Subject)
	newMessage.SetBody("text/html", mail.Body)
	if mail.Attachment != "" {
		newMessage.Attach(mail.Attachment)
	}

	dialer := gomail.NewDialer(smtp_server, smtp_port_int, smtp_username, smtp_password)

	if err := dialer.DialAndSend(newMessage); err != nil {
		// TODO: store into log table
		log.Println("Failed to send the email: " + err.Error())

		return err
	}

	return nil
}