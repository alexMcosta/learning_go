package main

import "fmt"

//Hello ... Go lint won't stop riding me about adding a comment here
func Hello() string {
	return "Hello, world"
}
func main() {
	fmt.Println(Hello())
}
