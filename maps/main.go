package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"Nicholas": "Hi!",
		"Kyle":     "cookie?",
	}

	myGreeting["Alex"] = "hello"
	myGreeting["Corey"] = "howdey"

	fmt.Println(myGreeting)
}
