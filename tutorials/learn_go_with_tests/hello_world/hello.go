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

	if language == spanish {
		return spanishPrefix + name
	}

	if language == french {
		return frenchPrefix + name
	}

	return englishPrefix + name
}
func main() {
	fmt.Println(Hello("World", ""))
}
