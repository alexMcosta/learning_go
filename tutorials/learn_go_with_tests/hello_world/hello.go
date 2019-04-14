package main

import "fmt"

const englishPrefix = "Hello, "

//Hello ... Go lint won't stop riding me about adding a comment here
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Spanish" {
		return "Hola, " + name
	}

	return englishPrefix + name
}
func main() {
	fmt.Println(Hello("World", ""))
}
