package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

type User struct {
	// gorm.Model
	ID    uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func initialMigration() {
	dsn := "root:password@tcp(127.0.0.1:3306)/gormtest?parseTime=true"
	db, err = gorm.Open("mysql", dsn)

	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.Close()

	db.AutoMigrate(&User{})
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dsn := "root:password@tcp(127.0.0.1:3306)/gormtest"
	db, err = gorm.Open("mysql", dsn)

	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.Close()

	var users []User
	db.Find(&users)

	json.NewEncoder(w).Encode(users)
}

func newUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dsn := "root:password@tcp(127.0.0.1:3306)/gormtest"
	db, err = gorm.Open("mysql", dsn)

	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.Close()

	// reqBody, _ := ioutil.ReadAll(r.Body)

	var user User
	// json.Unmarshal(reqBody, &user)
	json.NewDecoder(r.Body).Decode(&user)

	db.Create(&user)
	json.NewEncoder(w).Encode(user)
	// fmt.Fprintf(w, "new user created")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	dsn := "root:password@tcp(127.0.0.1:3306)/gormtest"
	db, err = gorm.Open("mysql", dsn)

	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	key := vars["id"]

	var user User
	db.Where("id = ?", key).Delete(&user)

	fmt.Fprintf(w, "Deleted User")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dsn := "root:password@tcp(127.0.0.1:3306)/gormtest"
	db, err = gorm.Open("mysql", dsn)

	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	key := vars["id"]

	// reqBody, _ := ioutil.ReadAll(r.Body)

	var user User
	db.Where("id = ?", key).Find(&user)

	json.NewDecoder(r.Body).Decode(&user)

	// json.Unmarshal(reqBody, &user)

	db.Save(&user)

	fmt.Fprint(w, "Update User")
}
