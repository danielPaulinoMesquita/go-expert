package main

import "fmt"

func soma(a, b *int) int {
	*a = 50 // <- * serve to get the value, at the moment, change the original value from variavel1

	return *a + *b
}

func main() {
	variavel1 := 10
	variavel2 := 30

	// passing the references
	result := soma(&variavel1, &variavel2)

	fmt.Println("New value of variavel1", variavel1)

	fmt.Println("Resultado:", result)

	pointers()

}
