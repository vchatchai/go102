package stnet

import (
	"crypto/tls"
	"net/smtp"
)

const EMAIL_SERVER = "smtp.gmail.com"

func SMTP(email, pass string) {
	auth := smtp.PlainAuth("", email, pass, EMAIL_SERVER)
	c, err := smtp.Dial("smtp.gmail.com:587")
	if err != nil {
		panic(err)
	}

	defer c.Close()
	config := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         EMAIL_SERVER,
	}

	if err = c.StartTLS(config); err != nil {
		panic(err)
	}

	if err = c.Auth(auth); err != nil {
		panic(err)
	}

	if err = c.Mail(email); err != nil {
		panic(err)
	}

	if err = c.Rcpt(email); err != nil {
		panic(err)
	}
	w, err := c.Data()
	if err != nil {
		panic(err)
	}

	msg := []byte("Hello this is content")
	if _, err = w.Write(msg); err != nil {
		panic(err)
	}

	err = w.Close()
	if err != nil {
		panic(err)
	}
	err = c.Quit()
	if err != nil {
		panic(err)
	}
}
