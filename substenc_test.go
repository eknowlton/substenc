package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncrypt(t *testing.T) {
	got := Encrypt("encrypt this")
	want := "qzodkbf ftue"

	assert.Equal(t, want, got)
}

func TestDecrypt(t *testing.T) {
	got := Decrypt("qzodkbf ftue")
	want := "encrypt this"

	assert.Equal(t, want, got)
}
