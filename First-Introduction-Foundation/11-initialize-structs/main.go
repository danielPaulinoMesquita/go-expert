package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	daniel := Cliente{
		Nome:  "Daniel",
		Idade: 30,
		Ativo: true,
	}

	fmt.Println(daniel)

	// you can change the value of Cliente. like this.
	daniel.Idade = 45

	fmt.Println(daniel.Nome)
	fmt.Println("--- Daniel all values after change ---")
	fmt.Println(daniel)
}
