package utils

import (
	"os"
	"testing"
)

func TestNewMail_msg_plain(t *testing.T) {
	from := "frodo@example.com"
	to := []string{"to@example.com"}
	subject := "Hello!"
	body := "Hello World!"
	mail := NewMail(from, to, subject, body)

	want := []byte("From: frodo@example.com\r\n" +
		"Subject: Hello!\r\n" +
		"MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\r\n" +
		"\r\n" +
		"Hello World!")

	got := mail.msg()

	if string(want) != string(got) {
		t.Errorf("got %s", string(got))
		t.Errorf("want %s", string(want))
	}
}

func TestNewMail_msg_html(t *testing.T) {
	from := "frodo@example.com"
	to := []string{"to@example.com"}
	subject := "Hello!"
	body := "<p>Hello World!</p>"
	mail := NewMail(from, to, subject, body)
	mail.ContentType = "text/html"

	want := []byte("From: frodo@example.com\r\n" +
		"Subject: Hello!\r\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n" +
		"\r\n" +
		"<p>Hello World!</p>")

	got := mail.msg()

	if string(want) != string(got) {
		t.Errorf("got %s", string(got))
		t.Errorf("want %s", string(want))
	}
}

func TestNewMail_server(t *testing.T) {
	defer os.Unsetenv("EMAIL_")
	os.Setenv("EMAIL_HOST", "mail")
	os.Setenv("EMAIL_PORT", "25")
	mail := NewMail("", []string{}, "", "")

	want := "mail:25"
	got := mail.server()

	if want != got {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestNewMail_auth_creds(t *testing.T) {
	defer os.Unsetenv("EMAIL_")
	os.Setenv("EMAIL_HOST", "mail")
	os.Setenv("EMAIL_USER", "user")
	os.Setenv("EMAIL_PASSWORD", "pass")
	mail := NewMail("", []string{}, "", "")

	if mail.auth() == nil {
		t.Error("expected auth not to be nil")
	}
}

func TestNewMail_auth_nocreds(t *testing.T) {
	defer os.Unsetenv("EMAIL_")
	os.Setenv("EMAIL_HOST", "mail")
	os.Setenv("EMAIL_USER", "")
	os.Setenv("EMAIL_PASSWORD", "")
	mail := NewMail("", []string{}, "", "")

	if mail.auth() != nil {
		t.Error("expected auth to be nil")
	}
}
