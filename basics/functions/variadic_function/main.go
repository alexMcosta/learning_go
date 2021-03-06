/***************************************************************
Variadic Function

This is an example from working on this:
https://www.udemy.com/googlego/learn/v4/content

I AM SUPER EXCITED ABOUT THIS! Almost to the point where this might be a hammmer(Refrence to everything being a nail).

Variadic functions seem (I will use seem about everything until I know it as a fact) that they are functions that can take more then one input by creating a slice and filling it by the given variable

below we have a variadic function that states
-------func add_grocery(a ...string){...}---------
What this is saying is that when called create a slice or array named 'a' where each value will be a string. even if there is one value it comes in as an array or slice.

***************************************************************/
package main

import "fmt"

var g_groceries []string

//The elipses allows multiple imputs
func add_grocery(a ...string) {
	for _, d := range a {
		g_groceries = append(g_groceries, d)
	}
}

func list_groceries() {
	fmt.Println("Grocery list as follows:")

	// A For range is built like this:
	// for INDEX, VALUE := range SLICE
	// _ lets us throw away the index
	for _, d := range g_groceries {
		fmt.Println(d)
	}

	//This is a variadic argument: example(slice...)
	//This means it will take a slice and insert each index through the function
}

func main() {
	add_grocery("bread", "pizza", "soda", "pees", "magic cards") //This is what the variadic function allows us to do
	add_grocery("candy")
	add_grocery("soap")
	add_grocery("milk")
	add_grocery("honey")
	add_grocery("party")
	add_grocery("phone")
	add_grocery("dog")
	list_groceries()

}
