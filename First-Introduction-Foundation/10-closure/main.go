package main

import "fmt"

func main() {

	// this likely function anonymous (Real name is Closure)
	total := func() int {
		return sum(2, 7, 9, 3)
	}

	fmt.Println("Total using CLOSURE(anonymous): ", total())
}

func sum(numbers ...int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}

	return total
}
