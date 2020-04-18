package utils

import (
	"time"

	"github.com/dchest/passwordreset"
)

func GetTokenHash(login string) ([]byte, error) {
	return []byte(login), nil
}

func NewToken(login string, dur time.Duration) (token string, err error) {
	hash, err := GetTokenHash(login)
	if err != nil {
		return
	}
	secret := DecodeEnvKey("PASSWORD_RESET_KEY")
	token = passwordreset.NewToken(login, dur, hash, secret)
	return
}

func VerifyToken(token string) (login string, err error) {
	secret := DecodeEnvKey("PASSWORD_RESET_KEY")
	login, err = passwordreset.VerifyToken(token, GetTokenHash, secret)
	return
}
