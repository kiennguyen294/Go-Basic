package main

import (
	"fmt"
)

// Channel

// 1. Unbuffered channel
// 2. Buffered channel
// 3. select
// 4. close channel

func main() {
	// Unbuffered channel
	ch := make(chan int)

	go func() {
		ch <- 100
	}()

	fmt.Println(<-ch)
	fmt.Println("done")

	// Buffered channel
	fmt.Println("========== Buffered channel =============")

	ch1 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2

	fmt.Println(<-ch1)
	fmt.Println(<-ch1)


	// select
	fmt.Println("======== Select ================")
	queue := make(chan int)
	done := make(chan bool)

	go func() {
		for i :=0; i<10; i++ {
			queue <- i
		}
		done <- true
	}()

	for  {
		select {
			case v := <-queue:
				fmt.Println(v)
			case <- done:
				fmt.Println("done")
				return
		}
	}
}