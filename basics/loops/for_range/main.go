/************************************************************
For Range

This is an example from working on this:
https://www.udemy.com/googlego/learn/v4/content

For ranges are the answer to going through an array and is
the answer for a foreach loop

************************************************************/

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
