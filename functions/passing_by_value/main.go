/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
Passing Values

This is basically just a demonstration that a function will not change the value of a variable which I feel is very common.

The only way it could be changed is if the variable is global
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/
package main

import "fmt"

func take_value(i int) {
	fmt.Println("In take_value function. i is:", i, "Located at:", &i)
	i = 0
	fmt.Println("In take_value function. i is reset and value of i is:", i, "located at:", &i, "\n")
}

func reset_value(i int) int {
	fmt.Print("In reset_value function. i comes in as: ", i, " Located at:", &i, "\n")
	i = 0
	fmt.Print("In reset_value function. i returns as: ", i, " Located at:", &i, "\n \n")
	return i
}

func main() {
	i := 15
	fmt.Println("In main function. i is:", i, "Located at:", &i, "\n")
	take_value(i)
	fmt.Println("In main function after take_value. i is:", i, "Located at:", &i, "\n")
	i = reset_value(i)
	fmt.Println("In main function after reset_value takes i and replces mains i value. i is:", i, "Located at:", &i, "\n")

}
