package sender

import gomail "gopkg.in/gomail.v2"

// https://pkg.go.dev/gopkg.in/gomail.v2
func SendEmail(feed Feed, senderEmail, senderPassword, senderSMTP string) {
	m := gomail.NewMessage()
	m.SetHeader("From", senderEmail)
	m.SetHeader("To", feed.Email)
	m.SetHeader("Subject", feed.Feedname+" - Personalnewsletter")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	// m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer(senderSMTP, 587, senderEmail, senderPassword)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}
