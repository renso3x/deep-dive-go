package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

type User struct {
	UserID uint   `json:"userId" gorm:"primary_key"`
	Name   string `json:"name"`
	Todos  []Todo `json:"todos" gorm:"foreignkey:UserID"`
}

type Todo struct {
	TodoID uint   `json:"todoId" gorm:"primary_key:autoincrement"`
	Name   string `json:"name"`
	UserId uint   `json: "-"`
}

func initDb() {
	fmt.Println("Initial Migration DB")

	dsn := "root:password@tcp(127.0.0.1:3306)/gormtodo?parseTime=true"

	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		panic("Unable to connect to database")
	}

	db.AutoMigrate(&Todo{}, &User{})
}
