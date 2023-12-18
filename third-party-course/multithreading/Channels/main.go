package main

import "fmt"

func main() {
	chanell := make(chan string) // Empty

	go func() {
		chanell <- "Olá Mundo!" // It is full
	}()

	// Thread 1
	msg := <-chanell // Channel empties
	fmt.Println("Res: ", msg)
}
