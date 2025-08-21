package utils

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func SendEmail(recipient, name, subject, message string) error {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	user := os.Getenv("SMTP_USER")
	password := os.Getenv("SMTP_PASSWORD")
	if host == "" || user == "" || password == "" {
		err := fmt.Errorf("SMTP environment variables are not set")
		log.Println(err)
		return err
	}

	addr := fmt.Sprintf("%s:%s", host, port)
	auth := smtp.PlainAuth("", user, password, host)
	to := []string{recipient}
	msg := []byte("To: " + recipient + "\r\n" + message + "\r\n")

	err := smtp.SendMail(addr, auth, user, to, msg)
	if err != nil {
		log.Printf("SMTP Error: %v", err)
		return err
	}
	return nil
}