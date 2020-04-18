package main

import (
	"encoding/base64"
	"flag"
	"fmt"

	"github.com/gorilla/securecookie"
)

// generate a base64 encoded random key
// this can be used as the session key env vars
func main() {
	length := flag.Int("length", 32, "key length")

	flag.Parse()

	key := securecookie.GenerateRandomKey(*length)
	encoded := base64.StdEncoding.EncodeToString(key)

	fmt.Println("Length:")
	fmt.Println(*length)
	fmt.Println("Key:")
	fmt.Println(encoded)
}
