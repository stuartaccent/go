package conf

import (
	"encoding/base64"
	"os"

	"github.com/gorilla/sessions"
)

func getKey(envVar string) []byte {
	key := os.Getenv(envVar)
	b, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		panic(err)
	}
	return b
}

var CookieStore = sessions.NewCookieStore(
	getKey("SESSION_AUTH_KEY"),
	getKey("SESSION_ENCRYPTION_KEY"),
)
