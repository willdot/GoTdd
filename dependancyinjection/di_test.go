package dependancyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "will")

	got := buffer.String()
	want := "Hello, will"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
