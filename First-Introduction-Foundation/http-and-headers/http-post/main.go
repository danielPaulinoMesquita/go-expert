package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"name": "daniel"}`))
	resp, err := c.Post("http://www.google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil) //<- copywriter going to get the value, that is resp.body and throw in another func, in this case os.Stdout

}
