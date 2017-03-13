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

	// _ lets us throw away the counter
	for _, d := range g_groceries {
		fmt.Println(d)
	}
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
