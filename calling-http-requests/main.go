package main

import (
	"io"
	"net/http"
)

func main() {

	println("1 - passo")
	req, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}

	println("2 - passo")
	res, err := io.ReadAll(req.Body) // <-- this is necessary to read the datas from stream in the body
	if err != nil {
		panic(err)
	}
	println(string(res))
	req.Body.Close() // <-- necessário fechar o stream

	testingDefer()
}
