package main

import "fmt"

var g_groceries []string

func add_grocery(a string) {
	g_groceries = append(g_groceries, a)
}

func list_groceries() {
	fmt.Println("Grocery list as follows:")
	for _, d := range g_groceries {
		fmt.Println(d)
	}
}

func main() {
	add_grocery("bread")
	add_grocery("pizza")
	add_grocery("soda")
	add_grocery("candy")
	add_grocery("soap")
	add_grocery("milk")
	add_grocery("honey")
	add_grocery("party")
	add_grocery("phone")
	add_grocery("dog")
	list_groceries()

}
