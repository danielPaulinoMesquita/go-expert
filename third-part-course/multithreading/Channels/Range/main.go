package main

import (
	"fmt"
)

// Thread 1
func main() {
	ch := make(chan int)
	go publish(ch)
	reader(ch)
}

func reader(ch chan int) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)
	//close() It's necessary to avoid deadlock, it could throw a deadlock,
	//because the reader function waits to read the channel, but the channel can be empty,
	//so, forever will wait and then throws deadlock
}
