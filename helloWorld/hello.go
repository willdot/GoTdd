package hello

import (
	"fmt"
	"io"
	"os"
)

const englishPrefix = "Hello "
const spanishHelloPrefix = "Hola "
const frenchHelloPrefix = "Bonjour "

func main() {
	fmt.Println(Hello("world", ""))

	Greet(os.Stdout, "Will")
}

// Greet prints hello + the name passed int
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// Hello returns hello
func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishPrefix
	}
	return
}
