package main

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

func main() {

	var (
		mail = "mail@mail.com"
		pwd  = "123456"
	)

	m := gomail.NewMessage()
	m.SetHeader("From", mail)
	m.SetHeader("To", "xxxx@mail.com")
	m.SetHeader("Subject", "Hello!")
	m.SetHeader("text/html", "<h1>HELLO TESTE</h1>")

	d := gomail.NewDialer("smtp.gmail.com", 587, mail, pwd)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}
