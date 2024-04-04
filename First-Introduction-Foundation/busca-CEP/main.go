package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// AddressFromCep I used the https://transform.tools/json-to-go to generate this struct
type AddressFromCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {

	// go run main.go 13330-230
	//var api = "http://viacep.com.br/ws/13330-230/json/"

	// index = '_' and value= 'url'
	for _, cep := range os.Args[1:] {
		api := "http://viacep.com.br/ws/" + cep + "/json/"
		req, err := http.Get(api)
		if err != nil {
			_, err := fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
			if err != nil {
				return
			}
		}

		// this it will serve to close stream in the final of for, it's good to put here, because avoid some problem of leak
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			_, err := fmt.Fprintf(os.Stderr, "Erro ao ler resposta(corpo da resposta): %v\n", err)
			if err != nil {
				return
			}
		}

		var data AddressFromCep
		err = json.Unmarshal(res, &data)
		if err != nil {
			_, err := fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
			if err != nil {
				return
			}
		}

		fmt.Println(data.Localidade)
		fmt.Println(data.Logradouro)
		file, err := os.Create("cidade.txt")
		if err != nil {
			_, err := fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
			if err != nil {
				return
			}
		}

		defer func(file *os.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)
		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s", data.Cep, data.Localidade, data.Uf))

		fmt.Println("Gerado arquivo com Sucesso")

	}
}
