package main

import "fmt"

const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "
const spanish = "Spanish"
const french = "French"

//Hello ... Go lint won't stop riding me about adding a comment here
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishPrefix

	switch language {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
