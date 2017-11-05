package main

import "fmt"
import "time"

func main() {

	go counting("This")
	go counting("Should")
	counting("Not")
	go counting("Come")
	go counting("Out")
	counting("Normally")
	fmt.Println("DONE")
}

func counting(a string) {
	for i := 0; i < 3; i++ {
		fmt.Println(a)
		time.Sleep(time.Second / 1)
	}
}