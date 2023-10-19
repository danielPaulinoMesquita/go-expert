package main

import "fmt"

func main() {
	fmt.Println("Função variádica, ex: somando: ", sum(1, 5, 10, 6))
}

// remind! loooks like varargs from Java
func sum(numbers ...int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}

	return total
}
