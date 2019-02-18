package dependancyinjection

import (
	"bytes"
	"fmt"
)

// Greet prints hello + the name passed in
func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
