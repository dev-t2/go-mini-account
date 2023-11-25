package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const addr = ":8080"

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello Go"))
	})

	router.NotFoundHandler = http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("Not Found"))
	})

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	fmt.Printf("Server running at http://localhost%s\n", addr)

	log.Fatal(http.ListenAndServe(addr, loggedRouter))
}