package main

import "net/http"

func main() {

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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hello World!"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
