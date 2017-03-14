/****************************************************************
Loops

Go only has one loop and that is for and I kinda like that

****************************************************************/
package main

import "fmt"
import "time"

func main() {

	for i := 10; i > 0; i-- {
		//Printf %v means to print the value of i
		fmt.Printf("%v \n", i)
		//wait half a second
		time.Sleep(time.Second / 2)
	}

	fmt.Println("BOOM!")
}
