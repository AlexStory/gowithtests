package walker

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalker(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with a non field string",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"struct with nested fields",
			Person{
				"Chris",
				Profile{
					33,
					"London",
				},
			},
			[]string{"Chris", "London"},
		},
		{
			"pointers to things",
			&Person{
				"Chris",
				Profile{
					33,
					"London",
				},
			},
			[]string{"Chris", "London"},
		},
		{
			"slices",
			[]Profile{
				{33, "London"},
				{32, "Nashville"},
			},
			[]string{"London", "Nashville"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "London"},
				{32, "Nashville"},
			},
			[]string{"London", "Nashville"},
		},
		{
			"maps",
			map[string]string{
				"Foo": "bar",
				"Baz": "boz",
			},
			[]string{"bar", "boz"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("want %v, but got %v", test.ExpectedCalls, got)
			}

		})
	}
}