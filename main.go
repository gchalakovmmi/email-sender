package main

import (
	"os"
	"log"
	"net/smtp"
)

func main() {
	auth := smtp.PlainAuth("", os.Getenv("SENDER_EMAIL"), os.Getenv("PASSWORD"), os.Getenv("SMTP_HOST"))

	msg := []byte("To: "+os.Getenv("RECIPIENT_EMAIL")+"\r\n" +
		"Subject: An example email...\r\n" +
		"\r\n" +
		"This is the email body!\r\n")

	err := smtp.SendMail(
		os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"), 
		auth, 
		"api", 
		[]string{os.Getenv("RECIPIENT_EMAIL")}, 
		msg)

	if err != nil {
		log.Fatal(err)
	}
}
