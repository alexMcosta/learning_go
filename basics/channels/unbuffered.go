package main

import (
	"fmt"
	"time"
)

func main() {

	//Making of a channel
	channel := make(chan int)

	//These two go routines are talking
	//to eachother via the channel
	//These are unbuffered channels (They block)
	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<-channel)
		}
	}()

	time.Sleep(time.Second)

}
