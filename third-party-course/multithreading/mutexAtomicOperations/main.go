package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var number uint64 = 0

func main() {
	m := sync.Mutex{}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		m.Lock()
		number++
		m.Unlock()
		time.Sleep(300 * time.Millisecond)
		writer.Write([]byte(fmt.Sprintf("Você é o visitante número %d", number)))
	})

	// if you use the ab to request 10000 times, and check call server in localhost:3000,
	//it will appear 9960 or something like this, therefore 40 customers were not attends

	http.ListenAndServe(":3000", nil)
}
