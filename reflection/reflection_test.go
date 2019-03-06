package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"will"},
			[]string{"will"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"will", "bath"},
			[]string{"will", "bath"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"will", 29},
			[]string{"will"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(inputString string) {
				got = append(got, inputString)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
