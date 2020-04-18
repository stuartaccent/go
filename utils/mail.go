package utils

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"text/template"
)

type mail struct {
	From        string
	To          []string
	Subject     string
	Body        string
	ContentType string
	Charset     string
}

// NewMail construct a new mail
//
// Example 1: send a plain text email in a goroutine
//     go func() {
//         from := "frodo@example.com"
//         to := []string{"sam@example.com"}
//         subject := "Hello!"
//         body := "Hello World!"
//         mail := utils.NewMail(from, to, subject, body)
//         if err := mail.SendMail(); err != nil {
//             fmt.Println(err)
//         }
//     }()
//
// Example 2: send an html email in a goroutine
//     go func() {
//         from := "frodo@example.com"
//         to := []string{"sam@example.com"}
//         subject := "Hello!"
//         body := ""
//         mail := utils.NewMail(from, to, subject, body)
//         mail.ContentType = "text/html"
//         data := struct {
//             Name string
//         }{
//             Name: "Sam"
//         }
//         if err := mail.ParseTemplate("email.html", data); err != nil {
//             if err := mail.SendMail(); err != nil {
//                 fmt.Println(err)
//             }
//         }
//     }()
func NewMail(From string, To []string, Subject string, Body string) *mail {
	return &mail{
		From:        From,
		To:          To,
		Subject:     Subject,
		Body:        Body,
		ContentType: "text/plain",
		Charset:     "UTF-8",
	}
}

func (m *mail) ParseTemplate(templateFile string, data interface{}) error {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	m.Body = buf.String()
	return nil
}

func (m *mail) msg() []byte {
	from := "From: " + m.From + "\r\n"
	subject := "Subject: " + m.Subject + "\r\n"
	mime := fmt.Sprintf("MIME-version: 1.0;\nContent-Type: %s; charset=\"%s\";\r\n", m.ContentType, m.Charset)
	return []byte(from + subject + mime + "\r\n" + m.Body)
}

func (m *mail) auth() (auth smtp.Auth) {
	host := os.Getenv("EMAIL_HOST")
	username := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PASSWORD")
	if username == "" || password == "" {
		return
	}
	auth = smtp.PlainAuth("", username, password, host)
	return
}

func (m *mail) server() string {
	host := os.Getenv("EMAIL_HOST")
	port := os.Getenv("EMAIL_PORT")
	return fmt.Sprintf("%s:%s", host, port)
}

func (m *mail) SendMail() error {
	msg := m.msg()
	server := m.server()
	auth := m.auth()
	if auth == nil {
		c, err := smtp.Dial(server)
		if err != nil {
			return err
		}
		defer c.Close()
		if err = c.Mail(m.From); err != nil {
			return err
		}
		for _, addr := range m.To {
			if err = c.Rcpt(addr); err != nil {
				return err
			}
		}
		w, err := c.Data()
		if err != nil {
			return err
		}
		_, err = w.Write(msg)
		if err != nil {
			return err
		}
		err = w.Close()
		if err != nil {
			return err
		}
		return c.Quit()
	}
	return smtp.SendMail(server, auth, m.From, m.To, msg)
}
