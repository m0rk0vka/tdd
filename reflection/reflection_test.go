package reflection

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

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string filed",
			struct {
				Name string
			}{
				"vova",
			},
			[]string{
				"vova",
			},
		},
		{
			"struct with two string fileds",
			struct {
				Name string
				City string
			}{
				"Vova",
				"Moscow",
			},
			[]string{
				"Vova",
				"Moscow",
			},
		},
		{
			"struct with not only string fields",
			struct {
				Name string
				Age  int
			}{
				"Vova",
				7,
			},
			[]string{
				"Vova",
			},
		},
		{
			"nested struct",
			Person{
				"Vova",
				Profile{
					23,
					"Moscow",
				},
			},
			[]string{
				"Vova",
				"Moscow",
			},
		},
		{
			"pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
