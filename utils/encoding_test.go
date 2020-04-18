package utils

import (
	"encoding/base64"
	"os"
	"testing"
)

func TestDecodeEnvKey(t *testing.T) {
	defer os.Unsetenv("TEST_KEY")

	want := []byte("my_key")
	key := base64.StdEncoding.EncodeToString(want)

	os.Setenv("TEST_KEY", key)

	got := DecodeEnvKey("TEST_KEY")

	if len(want) != len(got) {
		t.Error("not the same length")
	}
	for i, v := range want {
		if v != got[i] {
			t.Errorf("got %v, want %v", got[i], v)
		}
	}
}
