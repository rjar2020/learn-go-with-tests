package reflection

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWalk(t *testing.T) {
	assert := require.New(t)
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			assert.Equal(test.ExpectedCalls, got)
		})
	}
}
