package main

import "fmt"

const spanish = "Spanish"
const esperanto = "Esperanto"
const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const esperantoPrefix = "Saluton, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := greetingPrefix(language)

	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case esperanto:
		prefix = esperantoPrefix
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}