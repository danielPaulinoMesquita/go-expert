package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

// main is considered thread, the first thread
//
//	1
func main() {
	// thread 2
	go task("A")

	// thread 3, go indicates a thread
	go task("B")

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(15 * time.Second)

}
