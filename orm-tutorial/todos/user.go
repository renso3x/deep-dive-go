package main

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
)

func allUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	dsn := "root:password@tcp(127.0.0.1:3306)/gormtodo?parseTime=true"

	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		panic("Unable to connect to database")
	}

	defer db.Close()

	var users []User
	db.Preload("Todos").Find(&users)

	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	dsn := "root:password@tcp(127.0.0.1:3306)/gormtodo?parseTime=true"

	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		panic("Unable to connect to database")
	}

	defer db.Close()

	var user User
	json.NewDecoder(r.Body).Decode(&user)

	db.Create(&user)

	json.NewEncoder(w).Encode(user)
}
