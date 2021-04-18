package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGreet(t *testing.T) {

	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	require.Equal(t, want, got)
}

func BenchmarkGreet(b *testing.B) {
	buffer := bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		Greet(&buffer, "ImTestPerson")
	}
}

func ExampleGreet() {
	Greet(os.Stdout, "Tester")
	/* Output: Hello, Tester*/
}
