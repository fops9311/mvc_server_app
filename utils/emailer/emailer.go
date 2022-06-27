package emailer

import (
	"fmt"
	"net/smtp"
	"os"
)

var from string
var password string

func init() {
}

func Send(to []string, subject string, payload string) (err error) {

	from = os.Getenv("EMAIL_LOGIN")    //"from@gmail.com"
	password = os.Getenv("EMAIL_PASS") //"<Email Password>"
	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	var body []byte
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body = ([]byte)(fmt.Sprintf(
		"Subject: %s \n%s\n\n%s",
		subject,
		mimeHeaders,
		payload,
	))

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return
}
