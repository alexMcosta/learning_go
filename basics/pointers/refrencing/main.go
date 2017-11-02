/**************************************************************
Pointer refrencing
**************************************************************/

package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	// when the type has a '*' it means that the variables type is a pointer to said type
	var b *int = &a

	fmt.Println(b)

}
