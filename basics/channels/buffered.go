package main

import (
	"fmt"
)

func main() {

	channel := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		close(channel)
	}()

	for n := range channel {
		fmt.Println(n)
	}

}
