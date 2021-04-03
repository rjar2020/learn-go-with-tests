package maps

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearch(t *testing.T) {

	asserts := require.New(t)
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		asserts.Equal(want, got)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("Bazinga")
		asserts.Error(err)
		asserts.Equal(ErrNotFound, err)
	})
}

func TestAdd(t *testing.T) {
	asserts := require.New(t)

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)

		asserts.Nil(err)
		assertDefinition(asserts, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, definition)

		asserts.Error(err)
		asserts.Equal(ErrWordExists, err)
		assertDefinition(asserts, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	asserts := require.New(t)

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		asserts.Nil(err)
		assertDefinition(asserts, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		asserts.Error(err)
		asserts.Equal(ErrWordDoesNotExist, err)
	})
}

func TestDelete(t *testing.T) {
	asserts := require.New(t)
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)
	_, err := dictionary.Search(word)

	asserts.Error(err)
	asserts.Equal(ErrNotFound, err)
}

func assertDefinition(asserts *require.Assertions, dictionary Dictionary, word, definition string) {
	retrievedDefinition, err := dictionary.Search(word)
	asserts.Nil(err)
	asserts.Equal(definition, retrievedDefinition)
}

func BenchmarkDictionary_Add(b *testing.B) {
	word := "test"
	definition := "this is just a test"

	//I'm not happy with this, as it's also adding up the exec time of randString()
	b.Run("new word", func(b *testing.B) {
		dictionary := Dictionary{}
		for i := 0; i < b.N; i++ {
			testWord := randString()
			_ = dictionary.Add(testWord, testWord)
		}
	})

	b.Run("existing word", func(b *testing.B) {
		dictionary := Dictionary{}
		for i := 0; i < b.N; i++ {
			_ = dictionary.Add(word, definition)
		}
	})
}

func BenchmarkDictionary_Search(b *testing.B) {
	word := "test"
	dictionary := Dictionary{}

	b.Run("nonexistent word", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = dictionary.Search(word)
		}
	})

	_ = dictionary.Add(word, "this is just a test")

	b.Run("existing word", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = dictionary.Search(word)
		}
	})
}

func BenchmarkDictionary_Update(b *testing.B) {
	word := "test"
	definition := "this is just a test"
	dictionary := Dictionary{}

	b.Run("nonexistent word", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = dictionary.Update(word, definition)
		}
	})

	_ = dictionary.Add(word, definition)

	b.Run("existing word", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = dictionary.Update(word, definition)
		}
	})
}

func BenchmarkDictionary_Delete(b *testing.B) {
	word := "test"
	//Not that accurate either, as it's including the Add exec time
	b.Run("new word", func(b *testing.B) {
		dictionary := Dictionary{}
		for i := 0; i < b.N; i++ {
			_ = dictionary.Add(word, "this is just a test")
			dictionary.Delete(word)
		}
	})

	b.Run("existing word", func(b *testing.B) {
		dictionary := Dictionary{}
		for i := 0; i < b.N; i++ {
			dictionary.Delete(word)
		}
	})
}

func ExampleDictionary_Add() {
	dictionary := Dictionary{}
	word := "test"
	definition := "this is just a test"
	_ = dictionary.Add(word, definition)
	fmt.Printf("%s ", "NONE")
	err := dictionary.Add(word, definition)
	fmt.Printf("\"%s\"", err)
	/* Output: NONE "cannot add word because it already exists"*/
}

func ExampleDictionary_Search() {
	dictionary := Dictionary{}
	word := "test"
	definition := "this is just a test"
	_, err := dictionary.Search(word)
	fmt.Printf("\"%s\" ", err)
	_ = dictionary.Add(word, definition)
	retrievedDefinition, _ := dictionary.Search(word)
	fmt.Printf("\"%s\"", retrievedDefinition)
	/* Output: "could not find the word you were looking for" "this is just a test"*/
}

func ExampleDictionary_Update() {
	dictionary := Dictionary{}
	word := "test"
	definition := "this is just a test"
	newDefinition := "New test definition"
	err := dictionary.Update(word, newDefinition)
	fmt.Printf("\"%s\" ", err)
	_ = dictionary.Add(word, definition)
	_ = dictionary.Update(word, newDefinition)
	fmt.Printf("%s", "NONE")
	/* Output: "cannot update word because it does not exist" NONE*/
}

func ExampleDictionary_Delete() {
	dictionary := Dictionary{}
	word := "test"
	definition := "this is just a test"
	dictionary.Delete(word)
	fmt.Printf("%s ", "NONE")
	_ = dictionary.Add(word, definition)
	dictionary.Delete(word)
	fmt.Printf("%s ", "NONE")
	/* Output: NONE NONE*/
}

func randString() string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 100)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
