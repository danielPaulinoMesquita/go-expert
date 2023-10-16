package main

import (
	"fmt"
)

const a = "Hellow World"

// ID Creating a type of variable
type ID int

var (
	b bool = true
	c int
	d string
	e float64 = 1.2
	f ID      = 21
)

func main() {
	fmt.Println("Testing about this.")
	fmt.Printf("Check the type %T", e)
	fmt.Printf("\nCheck the value %v", e)
	fmt.Printf("\nCheck the type %T", f)
}
