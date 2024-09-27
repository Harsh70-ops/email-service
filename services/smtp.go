package services

import (
	"email_service/config"
	"fmt"
	"net/smtp"
)

func SendEmails(to, from, subject, body string) error {

	//setting up authentication
	auth := smtp.PlainAuth("", config.SMTPUser, config.SMTPPass, config.SMTPHost)

	// setting up email
	message := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s", from, to, subject, body)
	//this will create the string like 
	// From: your-email@example.com
	// To: recipient@example.com
	// Subject: Test Email
	
	// This is the email body.
	

	//sending email , sending back error if any!!

	addr := fmt.Sprintf("%s:%s", config.SMTPHost, config.SMTPPort)
	//this above function returns a format
	// config.SMTPHost = "smtp.example.com"
    // config.SMTPPort = "587"

	return smtp.SendMail(addr, auth,from,[]string{to}, []byte(message))
}
