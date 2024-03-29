package reflection

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

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
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"Nested fields",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Barcelona"},
			},
			[]string{"London", "Barcelona"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "London"},
				{34, "Barcelona"},
			},
			[]string{"London", "Barcelona"},
		},
		{
			"Maps",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			assert.ElementsMatch(test.ExpectedCalls, got)
		})
	}

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Katowice"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Katowice"}

		Walk(aChannel, func(input string) {
			got = append(got, input)
		})

		assert.Equal(want, got)
	})

	t.Run("with functions", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}

		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		Walk(aFunction, func(input string) {
			got = append(got, input)
		})

		assert.Equal(want, got)
	})
}

func ExampleWalk() {
	example := Person{
		"Chris",
		Profile{33, "London"},
	}

	var got []string

	Walk(example, func(input string) {
		got = append(got, input)
	})

	fmt.Printf("%s", got)

	/* Output: [Chris London]*/
}

func BenchmarkWalk(b *testing.B) {
	example := Person{
		"Chris",
		Profile{33, "London"},
	}

	for i := 0; i < b.N; i++ {
		Walk(example, func(input string) {})
	}
}
