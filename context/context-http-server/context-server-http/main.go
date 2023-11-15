package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request initialized")
	defer log.Println("Request done")
	select {
	case <-time.After(5 * time.Second):
		// Imprime no command line stdout
		log.Println("Request processada com sucesso")
		// Imprime  no brownser
		w.Write([]byte("Request processada com sucesso"))
	case <-ctx.Done():
		// Imprime no command line stdout
		log.Println("Request cancelled success")

		// Imprime  no brownser
		w.Write([]byte("Request cancelled success"))
		// Imprime no command line stdout
		http.Error(w, "Request cancelada pelo cliente", http.StatusRequestTimeout)
	}

}
