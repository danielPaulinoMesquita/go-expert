package main

import (
	"fmt"
)

func main() {
	numeros := []string{"um", "dois", "três", "quatro"}

	for i := 0; i < 10; i++ {
		println(i)
	}

	// this case k is the index and v is value
	for k, v := range numeros {
		fmt.Printf("\nKew: %d and Value: %v", k, v)
	}

	i := 0
	// looks while
	for i < 10 {
		println(i)
		i++
	}

	If()

	//for { // loop infinite
	//	println("Hellow World!!!")
	//}
}
