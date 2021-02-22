package ints

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdder(t *testing.T) {
	got := Add(2, 2)
	want := 4
	assert.Equal(t, want, got, "The result of the sum operation is incorrect")
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
