package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
)

// EmailCreds holds fun facts about our email account
type EmailCreds struct {
	Username string
	Password string
	Hostname string
	Port     string
	Auth     smtp.Auth
}

var emailUser *EmailCreds

// InitEmail Setup Email, if available
func InitEmail() error {
	emailUser = new(EmailCreds)

	if os.Getenv("EMAIL_USERNAME") != "" {
		emailUser.Username = os.Getenv("EMAIL_USERNAME")
	} else {
		emailUser = nil
		return fmt.Errorf("EMAIL_USERNAME unset disabling email support")
	}

	if os.Getenv("EMAIL_PASSWORD") != "" {
		emailUser.Password = os.Getenv("EMAIL_PASSWORD")
	} else {
		emailUser = nil
		return fmt.Errorf("EMAIL_PASSWORD unset disabling email support")
	}

	if os.Getenv("EMAIL_HOSTNAME") != "" {
		emailUser.Hostname = os.Getenv("EMAIL_HOSTNAME")
	} else {
		emailUser = nil
		return fmt.Errorf("EMAIL_HOSTNAME unset disabling email support")
	}

	if os.Getenv("EMAIL_PORT") != "" {
		emailUser.Port = os.Getenv("EMAIL_PORT")
	} else {
		emailUser.Port = "587"
	}

	emailUser.Auth = smtp.PlainAuth("",
		emailUser.Username,
		emailUser.Password,
		emailUser.Hostname,
	)

	return nil
}

// EmailVerification sends a verification email
func EmailVerification(recipient Domain) error {
	var err error
	var msgText bytes.Buffer
	var msgHTML bytes.Buffer

	TemplateText, err := template.
		New("verificationemail.txt").
		ParseFiles("tmpls/verificationemail.txt")
	if err != nil {
		return err
	}

	TemplateHTML, err := template.
		New("verificationemail.html").
		ParseFiles("tmpls/verificationemail.html")
	if err != nil {
		return err
	}

	err = TemplateText.Execute(&msgText, recipient)
	if err != nil {
		return err
	}
	err = TemplateHTML.Execute(&msgHTML, recipient)
	if err != nil {
		return err
	}

	return SendEmail(recipient.OwnerEmail,
		"Verify your domain",
		msgText,
		msgHTML,
	)

}

// SendEmail sends an email!
func SendEmail(recipient string,
	subject string,
	msgText bytes.Buffer,
	msgHTML bytes.Buffer,
) error {

	if emailUser == nil {
		return fmt.Errorf("Email is disabled")
	}
	log.Println("Sending email to", recipient)
	return smtp.SendMail(emailUser.Hostname+":"+emailUser.Port,
		emailUser.Auth,
		emailUser.Username,
		[]string{recipient},
		[]byte("To: "+recipient+"\r\n"+
			"From: Domain Glass <noreply@domain.glass>\r\n"+
			"Subject: "+subject+"\r\n"+
			"Content-Type: multipart/alternative;\r\n"+
			"	boundary=\"----=_Part_-1234792361_708108731.1459450691577\"\r\n"+
			"\r\n"+
			"------=_Part_-1234792361_708108731.1459450691577\r\n"+
			"Content-Type: text/plain\r\n"+
			"\r\n"+
			msgText.String()+"\r\n"+
			"------=_Part_-1234792361_708108731.1459450691577\r\n"+
			"Content-Type:text/html\r\n"+
			"\r\n"+
			msgHTML.String()+"\r\n",
		))
}
