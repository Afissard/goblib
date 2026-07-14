package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	shared "github.com/afissard/goblib/shared/api"
)

func main() {
	resp, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	var hello shared.HelloResponse

	if err := json.NewDecoder(resp.Body).Decode(&hello); err != nil {
		panic(err)
	}

	fmt.Println(hello.Message)
}
