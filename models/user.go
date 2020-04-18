package models

import (
	"app/utils"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email      string `gorm:"type:varchar(120);unique_index;not null"`
	FirstName  string `gorm:"type:varchar(150);not null"`
	LastName   string `gorm:"type:varchar(150);not null"`
	Password   string `gorm:"type:varchar(255);not null"`
	IsActive   bool   `gorm:"type:boolean;not null;default:true"`
	LoggedInAt time.Time
}

func (u *User) String() string {
	return u.Email
}

func (u *User) CanLogin() bool {
	return u.IsActive
}

func (u *User) CheckPassword(password string) bool {
	return utils.ComparePasswords(u.Password, []byte(password))
}

func (u *User) SendPasswordReset(r *http.Request) error {
	token, _ := utils.NewToken(u.Password, 12*time.Hour)
	from := os.Getenv("EMAIL_DEFAULT_FROM_ADDRESS")
	to := []string{u.Email}
	subject := "Password reset at " + r.Host
	body := ""
	mail := utils.NewMail(from, to, subject, body)
	data := struct {
		Token string
		Host  string
	}{
		Token: token,
		Host:  r.Host,
	}
	if err := mail.ParseTemplate("templates/email/password_reset.txt", data); err == nil {
		if err := mail.SendMail(); err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func (u *User) SetPassword(password string) {
	hash := utils.HashAndSalt([]byte(password))
	u.Password = string(hash)
}
