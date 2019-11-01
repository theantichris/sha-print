package main

import (
	"crypto/sha256"
	"testing"
)

func TestHash(t *testing.T) {
	t.Run("it gets SHA256 hash", func(t *testing.T) {
		got := getHash("x")
		want := sha256.Sum256([]byte("x"))

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
