package main

import (
	"fmt"
)

const (
	spanish = "Spanish"
	french  = "French"
	vietnam = "VietNam"
)

const (
	englishHelloPrefix = "Hello, "
	frenchHelloPrefix  = "Bonjour, "
	spanishHelloPrefix = "Hola, "
	vietnamHelloPrefix = "Xin chào, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case vietnam:
		prefix = vietnamHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", "english"))
}
