package main

import "fmt"

//Hello ... Go lint won't stop riding me about adding a comment here
func Hello(name string) string {
	return "Hello, " + name
}
func main() {
	fmt.Println(Hello("World"))
}
