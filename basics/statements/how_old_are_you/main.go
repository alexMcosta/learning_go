package main

import "fmt"

func main() {

	age := 60

	if age < 13 {
		fmt.Println("You are a kid")
	} else if age < 20 {
		fmt.Println("You are a teenager")
	} else if age < 30 {
		fmt.Println("You are in your 20's")
	} else if age < 40 {
		fmt.Println("You are in your 30's")
	} else if age < 50 {
		fmt.Println("You are in your 40's")
	} else {
		fmt.Println("You are over the hill!")
	}

}
