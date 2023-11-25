package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type User struct {
	Id       int
	Nickname string
}

type Users struct {
	users []User
}

type CreateUserBody struct {
	Nickname string
}

var users = make([]User, 0)

const addr = ":8080"

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")

		json.NewEncoder(res).Encode(Users{users})
	}).Methods("GET")

	router.HandleFunc("/users", func(res http.ResponseWriter, req *http.Request) {
		var body CreateUserBody

		err := json.NewDecoder(req.Body).Decode(&body)

		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("Bad Request"))
		}
	}).Methods("POST")

	router.HandleFunc("/users", func(res http.ResponseWriter, req *http.Request) {
		// 
	}).Methods("PUT")

	router.HandleFunc("/users", func(res http.ResponseWriter, req *http.Request) {
		// 
	}).Methods("DELETE")

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	fmt.Printf("Server running at http://localhost%s\n", addr)

	log.Fatal(http.ListenAndServe(addr, loggedRouter))
}