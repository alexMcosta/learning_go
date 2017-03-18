package main

import "fmt"

func main() {
	myGreeting := map[string]string{}

	myGreeting["Alex"] = "hello"
	myGreeting["Corey"] = "howdey"

	fmt.Println(myGreeting)
}
