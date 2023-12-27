package main

import (
	"fmt"
	"sync"
)

// Thread 1
func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)
	go publish(ch)
	go reader(ch, &wg)
	wg.Wait()
}

func reader(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
		wg.Done()
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
