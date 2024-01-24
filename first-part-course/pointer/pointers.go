package main

import "fmt"

func pointers() {

	// todo '&' <- serve to get the value from his address
	var text string = "teste"
	var testeponteiros *string = &text
	b := testeponteiros // <- here, it receives the same address of testeponteiros

	fmt.Println(b)    // printa the address
	fmt.Println(text) // printa o valor

	*b = "novo valor"

	fmt.Println(text) // printa b o valor

	c := b
	*c = "new teste in letter c"

	fmt.Println(c)
	fmt.Println(text) // printa c o valor
}
