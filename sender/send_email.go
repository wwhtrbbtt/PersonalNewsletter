package sender

import gomail "gopkg.in/gomail.v2"

// https://pkg.go.dev/gopkg.in/gomail.v2
func SendEmail(feed Feed, template, senderEmail, senderPassword, senderSMTP, username string) {
	m := gomail.NewMessage()

	// Set all required headers
	m.SetHeader("From", senderEmail)
	m.SetHeader("To", feed.Email)
	m.SetHeader("Subject", feed.Feedname+" - Personalnewsletter")
	m.SetBody("text/html", feed.GetHTML(template)) // GetHTML converts the feed struct into HTML

	// Authenticate
	d := gomail.NewDialer(senderSMTP, 587, username, senderPassword)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}
