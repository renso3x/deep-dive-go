package main

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func getTodos(w http.ResponseWriter, r *http.Request) {
	dsn := "root:password@tcp(127.0.0.1:3306)/gormtodo?parseTime=true"

	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		panic("Unable to connect to database")
	}

	defer db.Close()

	var todos []Todo
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(todos)
}

func postTodos(w http.ResponseWriter, r *http.Request) {
	dsn := "root:password@tcp(127.0.0.1:3306)/gormtodo?parseTime=true"

	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		panic("Unable to connect to database")
	}

	defer db.Close()

	var t Todo
	err = json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Do something with the Person struct...
	// fmt.Fprintf(w, "Person: %+v", t)
	// json.NewDecoder(r.Body).Decode(&todo)
	db.Create(&t)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
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
