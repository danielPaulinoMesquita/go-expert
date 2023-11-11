package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	// if request delay for a long time, more than microsecond, it will throw error
	c := http.Client{Timeout: time.Microsecond}
	resp, err := c.Get("http://google.com")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	println(string(body))

}
