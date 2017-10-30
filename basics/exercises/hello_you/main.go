package main

import "fmt"

func main() {
	fmt.Println("What is your name?")
	var name string
	fmt.Scan(&name)
	fmt.Println("Hello", name)
}
