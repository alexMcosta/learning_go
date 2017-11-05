package main

import "fmt"

func main() {

	channel := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		done <- true
	}()

	go func() {
		for i := 10; i > 0; i-- {
			channel <- i
		}
		done <- true
	}()

	go func() {
		<-done
		<-done
		close(channel)
	}()

	for n := range channel {
		fmt.Println(n)
	}

}
