package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Learning about golang")
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Println("Error while fetching the response", err)
		return
	}

	data, erro := io.ReadAll(res.Body)
	if erro != nil {
		fmt.Println("error while parsing data", erro)
		return
	}
	fmt.Println("the return value is", string(data))
}
