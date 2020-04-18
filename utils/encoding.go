package utils

import (
	"encoding/base64"
	"os"
)

func DecodeEnvKey(envVar string) []byte {
	key := os.Getenv(envVar)
	b, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		panic(err)
	}
	return b
}
