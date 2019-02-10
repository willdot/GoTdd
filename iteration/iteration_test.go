package iteration

import (
	"fmt"
	"testing"
)

func TestRepeatFive(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	assert(t, expected, repeated)
}

func TestRepeatTen(t *testing.T) {
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	assert(t, expected, repeated)
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func assert(t *testing.T, expected string, actual string) {
	if actual != expected {
		t.Errorf("expected '%s' but got '%s'", expected, actual)
	}
}
