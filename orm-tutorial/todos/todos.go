package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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
	db.Find(&todos)
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
	db.Create(&t)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	dsn := "root:password@tcp(127.0.0.1:3306)/gormtodo?parseTime=true"

	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		panic("Unable to connect to database")
	}

	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	var t Todo
	db.Where("todo_id = ?", id).Find(&t)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(t)

}

func updateTodos(w http.ResponseWriter, r *http.Request) {
	dsn := "root:password@tcp(127.0.0.1:3306)/gormtodo?parseTime=true"

	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		panic("Unable to connect to database")
	}

	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	var t Todo
	db.Where("todo_id = ?", id).Find(&t)

	json.NewDecoder(r.Body).Decode(&t)

	db.Save(&t)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
}

func deleteTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dsn := "root:password@tcp(127.0.0.1:3306)/gormtodo?parseTime=true"

	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		panic("Unable to connect to database")
	}

	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	var t Todo
	db.Where("todo_id = ?", id).Delete(&t)

	fmt.Fprintf(w, "Deleted successfully")
}
