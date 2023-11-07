package main

import (
	"net/http"
)

func main() {
	//todo check later, why this doesn't work 'http.Dir'
	fileServer := http.FileServer(http.Dir("./public/index.html"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("HELLO FROM BLOG!"))
	})

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		return
	}
}
