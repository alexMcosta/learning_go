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
		//When a channel is closed it can not take anymore input
		//But it can still send its data.
		close(channel)
	}()

	for n := range channel {
		fmt.Println(n)
	}

}
