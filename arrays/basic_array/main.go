package main

import "fmt"

const g_cap int = 5 // Total compacity of grocery list

var g_groceries [g_cap]string
var g_len int = 0

func add_grocery(a string) {
	if g_len < g_cap {
		g_groceries[g_len] = a
		g_len++
	} else {
		fmt.Println("Too many groceries")
	}
}

func list_groceries() {
	fmt.Println("Grocery list as follows:")
	for i := 0; i < g_len; i++ {
		fmt.Println(g_groceries[i])
	}
}

func main() {
	add_grocery("bread")
	add_grocery("pizza")
	add_grocery("soda")
	add_grocery("candy")
	add_grocery("soap")
	add_grocery("milk")
	list_groceries()

}
