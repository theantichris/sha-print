package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"testing"
)

func TestHash(t *testing.T) {
	t.Run("it gets SHA256 hash", func(t *testing.T) {
		got := get256("x")
		want := sha256.Sum256([]byte("x"))

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("it gets SHA384 hash", func(t *testing.T) {
		got := get384("x")
		want := sha512.Sum384([]byte("x"))

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("it gets SHA512 hash", func(t *testing.T) {
		got := get512("x")
		want := sha512.Sum512([]byte("x"))

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
