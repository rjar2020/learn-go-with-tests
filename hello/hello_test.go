package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assert.Equal(t, want, got, "The greetings message doesn't contain the right information")
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assert.Equal(t, want, got, "The greetings message when no name and language supplied doesn't contain the right information")
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assert.Equal(t, want, got, "The greetings message in spanish doesn't contain the right information")
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("James", "French")
		want := "Bonjour, James"
		assert.Equal(t, want, got, "The greetings message in french doesn't contain the right information")
	})
}
