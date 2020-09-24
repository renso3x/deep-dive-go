package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func handleRequest() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", helloWorld).Methods("GET")
	r.HandleFunc("/users", allUsers).Methods("GET")
	r.HandleFunc("/users", newUser).Methods("POST")
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	r.HandleFunc("/users/{id}", updateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8081", r))
}

func main() {
	fmt.Println("Go is Running")
	initialMigration()
	handleRequest()
}
