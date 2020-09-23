package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string "json:'name'"
}

func main() {
	fmt.Println("Go MySql")

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	var user User

	// QUery a single Row

	err = db.QueryRow("SELECT * from users where name=?", "ROMEO").Scan(&user.Name)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(user.Name)

	// SELECT ALL
	// results, err := db.Query("SELECT * from users")

	// if err != nil {
	// 	panic(err.Error())
	// }

	// for results.Next() {
	// 	var user User
	// 	// for each row, scan the result into our tag composite object

	// 	err := results.Scan(&user.Name)

	// 	if err != nil {
	// 		panic(err.Error())
	// 	}

	// 	fmt.Println(user.Name)
	// }
	// defer results.Close()

}

// func insertToDb() {
// 	insert, err := db.Query("INSERT INTO users values ('ROMEO')")

// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer insert.Close()

// 	fmt.Println("Succesfully inserted to mysql db.")
// }
