package main

import "fmt"

func testingDefer() {
	fmt.Println("Primeira segunda")
	defer fmt.Println("Segunda segunda") // Defer will hold this line for the final
	fmt.Println("Terceira segunda")
}
