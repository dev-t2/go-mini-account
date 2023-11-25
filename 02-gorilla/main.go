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

	loggedRouter := handlers.CombinedLoggingHandler(os.Stdout, router)
	
	recoveredRouter := handlers.RecoveryHandler()

	fmt.Printf("Server running at http://localhost%s\n", addr)

	log.Fatal(http.ListenAndServe(addr, recoveredRouter(loggedRouter)))
}