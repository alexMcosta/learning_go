/**************************************************************
Pointer derefrencing
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

	//By applying the asterisk to 'b' it is stating "Give me the value of 'b' and not the location"
	fmt.Println(*b) // Derefrencing

}
