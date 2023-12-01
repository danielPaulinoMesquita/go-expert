package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int
	Saldo  int
}

// Conta2 example for using tags
// including, you can put some validates for a property
type Conta2 struct {
	Numero int `json:"numero"`
	Saldo  int `json:"saldo"`
	//Salary  int `json:"-"` <-- it will ignore the Salary because the '-'
}

func main() {

	conta := Conta2{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta) // Marshal save the result that is the json

	if err != nil {
		println(err)
	}

	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta) // Encode doesn't save the json, so you should print or write
	if err != nil {
		println(err)
	}

	jsonPure := []byte(`{"numero": 2, "saldo": 200}`)
	var contaX Conta
	err = json.Unmarshal(jsonPure, &contaX) // <== for deserializing , you need a struct,
	// and because this you pass the reference memory of struct Conta, and one more thing,
	// at the moment of parsing, ots necessary the struct has the same properties from json
	// So, if this json was like this {"n":2, "s":200} it will receive error

	if err != nil {
		println(err)
	}

	println(contaX.Saldo)

}
