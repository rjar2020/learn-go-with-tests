package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHello(t *testing.T) {

	assert := require.New(t)

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assert.Equal(want, got, "The greetings message doesn't contain the right information")
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assert.Equal(want, got, "The greetings message when no name and language supplied doesn't contain the right information")
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assert.Equal(want, got, "The greetings message in spanish doesn't contain the right information")
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("James", "French")
		want := "Bonjour, James"
		assert.Equal(want, got, "The greetings message in french doesn't contain the right information")
	})
}
