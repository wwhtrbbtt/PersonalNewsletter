package sender

import (
	"fmt"
	"net/smtp"
)

func SendEmail(feed Feed, senderEmail, senderPassword, senderSMTP string) {
	auth := smtp.PlainAuth("", senderEmail, senderPassword, senderSMTP)

	subject := fmt.Sprintf("%s - PersonalNewsletter\n", feed.Feedname)
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := fmt.Sprintf("<html><body><h1>Hello, %s</h1></body></html>", feed.Greetingname)
	msg := []byte(subject + mime + body)

	smtp.SendMail(senderSMTP+":587", auth, senderEmail, []string{feed.Email}, msg)

}
