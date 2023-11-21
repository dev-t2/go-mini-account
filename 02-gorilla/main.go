package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const addr = ":8080"

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello Go"))
	})

	fmt.Printf("Server running at http://localhost%s\n", addr)

	log.Fatal(http.ListenAndServe(addr, router))
}