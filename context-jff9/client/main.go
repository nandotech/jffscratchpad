package main

import (
	"os"
	"io"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)

	}
	io.Copy(os.Stdout, res.Body)
}