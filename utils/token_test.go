package utils

import (
	"testing"
	"time"
)

func TestGetTokenHash(t *testing.T) {
	token := "my_token"
	want := []byte(token)
	got, err := GetTokenHash(token)
	if err != nil {
		t.Error(err)
	}
	if len(want) != len(got) {
		t.Error("not the same length")
	}
	for i, v := range want {
		if v != got[i] {
			t.Errorf("got %v, want %v", got[i], v)
		}
	}
}

func TestNewToken(t *testing.T) {
	token := "my_token"
	if _, err := NewToken(token, 10*time.Second); err != nil {
		t.Error(err)
	}
}

func TestVerifyToken(t *testing.T) {
	token := "my_token"
	bs64, _ := NewToken(token, 10*time.Second)

	decoded, err := VerifyToken(bs64)
	if err != nil {
		t.Error(err)
	}
	if token != decoded {
		t.Errorf("got %s, want %s", decoded, token)
	}
}
