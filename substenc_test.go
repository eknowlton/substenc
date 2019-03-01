package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncrypt(t *testing.T) {
	got := Encrypt("encrypt this")
	want := "qzoDKBF FtuE"

	assert.Equal(t, want, got)
}

func TestDecrypt(t *testing.T) {
	got := Decrypt("qzoDKBF FtuE")
	want := "encrypt this"

	assert.Equal(t, want, got)
}
