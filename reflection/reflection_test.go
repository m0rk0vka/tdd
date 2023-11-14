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
		{
			"slices",
			[]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
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
	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}
		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})
		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})
	t.Run("with chans", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{23, "Moscow"}
			aChannel <- Profile{80, "London"}
			close(aChannel)
		}()

		var got []string
		var want = []string{
			"Moscow",
			"London",
		}
		walk(aChannel, func(value string) {
			got = append(got, value)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
