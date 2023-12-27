package main

import "fmt"

func receive(nome string, hello chan<- string) {
	hello <- nome
	// The hello channel only will receive string, named Receive only
}

func read(data <-chan string) {
	fmt.Println(<-data)
	// The arrow on the left means emptying of channel, named Send only
}

func main() {
	hello := make(chan string)
	go receive("Hello", hello)
	read(hello)
}
