package main

import "fmt"

const delimiter = ", "
const englishHelloPrefix = "Hello"
const spanishHelloPrefix = "Hola"
const frenchHelloPrefix = "Bonjour"
const languageSpanish = "Spanish"
const languageFrench = "French"

const defaultName = "World"

func Hello(name string, language string) string {
	if name == "" {
		name = defaultName
	}

	return greetingPrefix(language) + delimiter + name
}

func greetingPrefix(language string) (helloPrefix string) {
	switch language {
	case languageSpanish:
		helloPrefix = spanishHelloPrefix
	case languageFrench:
		helloPrefix = frenchHelloPrefix
	default:
		helloPrefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Bob", ""))
}
