package main

import (
	"fmt"
	"log"
	"net/http"
)

const addr = ":8080"

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello Go"))
	})

	fmt.Printf("Server running at http://localhost%s\n", addr)

	log.Fatal(http.ListenAndServe(addr, nil))
}