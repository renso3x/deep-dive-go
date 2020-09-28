package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequest() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/todos", getTodos).Methods("GET")
	r.HandleFunc("/todos", postTodos).Methods("POST")
	r.HandleFunc("/todo/{id}", getTodo).Methods("GET")
	r.HandleFunc("/todo/{id}", updateTodos).Methods("PUT")
	r.HandleFunc("/todo/{id}", deleteTodos).Methods("DELETE")

	r.HandleFunc("/users", allUsers).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", r))
}

func main() {
	fmt.Println("UserTodo ORM Rest API")
	initDb()
	handleRequest()
}
