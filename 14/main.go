package main

import "fmt"

// Any class(struct) that has the method Desativar()
// automatic implements this interface,
// Interfaces don't use variables in his body
type Pessoa interface {
	// Salary string this is not possible
	Desativar()
	Ativar(identy string)
}

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

// TODO STRUCT(CLASSES)
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
	//nomeDaVariavel Endereco <-  can be declared like this
}

type Empresa struct {
	Nome string
}

// this method is implementing on backstage the interface Pessoa
func (empresa Empresa) Desativar() {
	fmt.Println("Implemented By EMPRESA")
}
func (empresa Empresa) Ativar(string2 string) {
	fmt.Println("ATIVAR - Implemented By EMPRESA")
}

// Create a method  of Cliente, now on backstage it's implementing interface Pessoa
func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("The customer %s foi desativado: ", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func Ativar(pessoa Pessoa) {
	pessoa.Ativar("ativando:")
}

func main() {
	daniel := Cliente{
		Nome:  "Daniel",
		Idade: 30,
		Ativo: true,
	}

	empresaMinha := Empresa{
		Nome: "Teste",
	}

	empresaMinha.Desativar() // <-- This works
	// Desativacao(empresaMinha) // <-- this as welL only if Empresa has one method
	Ativar(empresaMinha) // <-- this as welL only if Empresa has one method

	// now, we are using the method
	daniel.Desativar()

	// Desativacao(daniel) <-- this looks like inheritance, so if Desativacao has two methods, Cliente has to implement these methods too.
}
