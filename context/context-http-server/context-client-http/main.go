package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	// time of delay to receive the response from server
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Here is only for creating the context of resquest
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}

	// here is to do the request create recently
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	// defer to close the request
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
