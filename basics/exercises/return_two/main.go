package main

import "fmt"

func main() {

	x, y := return_two(5)
	fmt.Println(x, y)
}

func return_two(a int) (int, bool) {

	return a / 2, a%2 == 0

}
