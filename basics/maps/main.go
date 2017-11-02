/**********************************************************
Maps are key vlue refrence types

**********************************************************/

package main

import "fmt"

func main() {
	myGreeting := map[string]string{} //Initiate a blank map

	myGreeting["Alex"] = "hello"
	myGreeting["Corey"] = "howdey"

	//fmt.Println(myGreeting)
	//fmt.Println(myGreeting["Alex"])

	//Using a For Range with a map
	for v, k := range myGreeting {
		fmt.Println(v, k)
	}

	// Deleteing from a map is like so:
	delete(myGreeting, "Corey")
	//fmt.Println(myGreeting)

	//Nested Maps
	myFamily := map[string]map[string]int{
		"Nat": map[string]int{
			"Daughters": 2,
			"Sons":      0,
			"Pets":      1,
		},
		"Ashley": map[string]int{
			"Daughters": 0,
			"Sons":      1,
			"Pets":      2,
		},
		"Holishia": map[string]int{
			"Daughters": 0,
			"Sons":      1,
			"Pets":      12,
		},
		"Teddy": map[string]int{
			"Daughters": 1,
			"Sons":      0,
			"Pets":      0,
		},
		"Kevin": map[string]int{
			"Daughters": 0,
			"Sons":      1,
			"Pets":      0,
		},
		"Denice": map[string]int{
			"Daughters": 0,
			"Sons":      0,
			"Pets":      3,
		},
	}

	//fmt.Println(myFamily)
	//fmt.Println(myFamily["Teddy"])
	fmt.Println(myFamily["Teddy"]["Sons"])

}
