/**********************************************************
Maps are key vlue refrence types

**********************************************************/

package main

import "fmt"

func main() {
	myGreeting := map[string]string{} //Initiate a blank map

	myGreeting["Alex"] = "hello"
	myGreeting["Corey"] = "howdey"

	fmt.Println(myGreeting)
	fmt.Println(myGreeting["Alex"])

	//Using a For Range with a map
	for v, k := range myGreeting {
		fmt.Println(v, k)
	}

	// Deleteing from a map is like so:
	delete(myGreeting, "Corey")
	fmt.Println(myGreeting)
}
