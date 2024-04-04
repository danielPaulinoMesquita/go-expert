package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// to create a file
	f, err := os.Create("arquivo.txt")
	if err != nil { // if is not null means that happened some error!
		panic(err)
	}

	// to write in the file 'f'
	//tamanho, err := f.WriteString("Escrevendo alguma coisa aqui!") <-- one way to write
	tamanho, err := f.Write([]byte("Escrevendo alguma coisa aqui para realizar testes!"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho %d bytes \n", tamanho)

	err = f.Close()
	if err != nil {
		return
	}

	// leitura do arquivo
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
		//fmt.Printf("%s", string(buffer[:n]))
	}

	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}

}
