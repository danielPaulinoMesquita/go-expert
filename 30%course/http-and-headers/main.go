package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

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

	// go run  http-and-headers/*
	// then, on the web browser write the url above
	//http://localhost:8080/?cep=12313123
	// Looks like an endpoint in Go
	http.HandleFunc("/", BuscarCep)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func BuscarCep(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return // don't forget this
	}

	cepParam := r.URL.Query().Get("cep")
	println("result cep path: ", cepParam)
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return // don't forget this. return
	}
	// Adding header to response of BuscarCep
	cep, err := BuscaCep(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// one way to response with json
	//result, err := json.Marshal(cep)
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//w.Write(result)

	json.NewEncoder(w).Encode(cep)

}

func BuscaCep(cep string) (*AddressFromCep, error) {
	resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var c AddressFromCep
	err = json.Unmarshal(body, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
