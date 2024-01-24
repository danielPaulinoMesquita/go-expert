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

// Create a method  of Cliente
func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("The customer %s foi desativado: ", c.Nome)
}

func main() {
	daniel := Cliente{
		Nome:  "Daniel",
		Idade: 30,
		Ativo: true,
	}

	// now, we are using the method
	daniel.Desativar()
}
