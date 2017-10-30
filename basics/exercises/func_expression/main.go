package main

import "fmt"

func main() {

	functional_expression := func(a int) (int, bool) {
		return a / 2, a%2 == 0
	}
	fmt.Println(functional_expression(2))
}
