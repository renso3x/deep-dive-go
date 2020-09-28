package main

import (
	"encoding/json"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var todos []Todo
	db.Find(&todos)
	// fmt.Fprintf(w, "Hello World")

	json.NewEncoder(w).Encode(todos)
}

func postTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func getTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func updateTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func deleteTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}
