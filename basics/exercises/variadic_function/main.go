package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 3, 0}

	greatest(a...)
	greatest(123, 34, 23, 4, 5, 6)
}

func greatest(a ...int) {
	var b int
	for _, i := range a {
		if i > b {
			b = i
		}
	}
	fmt.Println(b)
}
