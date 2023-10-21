package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
	//nomeDaVariavel Endereco <-  can be declared like this
}

func main() {
	daniel := Cliente{
		Nome:  "Daniel",
		Idade: 30,
		Ativo: true,
	}

	// It can understand that Endereco make part of the Cliente
	daniel.Logradouro = "Rua tal"
	daniel.Numero = 21
	daniel.Cidade = "Recanto"
	daniel.Estado = "DF"

	// You can pass the value like this as well
	daniel.Endereco.Numero = 22

	fmt.Println(daniel)
	fmt.Println(daniel.Nome)
	fmt.Println("--- Daniel all values after change ---")
	fmt.Println(daniel)
}
