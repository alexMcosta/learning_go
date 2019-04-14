package main

import "fmt"

const englishPrefix = "Hello, "

//Hello ... Go lint won't stop riding me about adding a comment here
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishPrefix + name
}
func main() {
	fmt.Println(Hello("World"))
}
