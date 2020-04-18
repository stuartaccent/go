package utils

import (
	"testing"
)

func TestHashAndSalt(t *testing.T) {
	pw := "password"
	hash := HashAndSalt([]byte(pw))

	if hash == "" {
		t.Error("not expecting a blank string")
	}
}
func TestComparePasswords(t *testing.T) {
	pw := "password"
	hash := HashAndSalt([]byte(pw))
	if valid := ComparePasswords(hash, []byte(pw)); !valid {
		t.Errorf("got %t, want %t", valid, true)
	}
	if valid := ComparePasswords(hash, []byte("Password")); valid {
		t.Errorf("got %t, want %t", valid, false)
	}
	if valid := ComparePasswords(hash, []byte("")); valid {
		t.Errorf("got %t, want %t", valid, false)
	}
}
