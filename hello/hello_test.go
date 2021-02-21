package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHello(t *testing.T) {

	assert := require.New(t)

	t.Run("say hello to Chris", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assert.Equal(want, got)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assert.Equal(want, got)
	})

	t.Run("say hello in Spanish (Hola)", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assert.Equal(want, got)
	})

	t.Run("say hello in French (Bonjour)", func(t *testing.T) {
		got := Hello("James", "French")
		want := "Bonjour, James"
		assert.Equal(want, got)
	})
}
