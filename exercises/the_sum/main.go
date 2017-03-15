package main

import "fmt"

func main() {

	fmt.Println("Give me a small number and a larger number:")

	var a, b, remainder, bigger, smaller int

	fmt.Scan(&a, &b)

	if a > b {
		remainder = a % b
		bigger = a
		smaller = b
	} else if b > a {
		remainder = b % a
		bigger = b
		smaller = a
	} else {
		fmt.Println("Those are the same number!")
		return
	}

	fmt.Println("The remainder of ", bigger, " devided by ", smaller, " is ", remainder)
}
