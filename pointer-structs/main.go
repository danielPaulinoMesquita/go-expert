package main

import "fmt"

type Conta struct {
	saldo int
}

// NewConta this is a usual practice to point the real Conta
// looks like a singleton
func NewConta() *Conta {
	return &Conta{saldo: 0}
}

// passing the real value Conta, without '*' it only a copy
func (c *Conta) simulador(valor int) int {
	c.saldo += valor
	return c.saldo
}

func main() {

	conta := Conta{saldo: 100}
	conta.simulador(20)
	fmt.Println(conta.simulador(30))

	println(conta.saldo)

}
