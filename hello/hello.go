package hello

import "fmt"

const helloPrefix = "hello, "
const spanishHelloPrefix = "hola, "
const frenchHelloPrefix = "bonjour, "

func hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) string {
	prefix := helloPrefix
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	}
	return prefix
}

func main() {
	fmt.Println(hello("stas", ""))
}
