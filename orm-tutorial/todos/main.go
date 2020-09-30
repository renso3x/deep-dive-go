package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var mySigningKey = []byte("thisissupersecret")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "Romeo Enso"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func NewToken(w http.ResponseWriter, r *http.Request) {
	token, err := GenerateJWT()

	if err != nil {
		fmt.Fprintf(w, "Something went wrong: %s", err.Error())
	}

	fmt.Fprintf(w, string(token))
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {

			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {

				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error")
				}

				return mySigningKey, nil

			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}

		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

func handleRequest() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/generate-token", NewToken).Methods("GET")
	r.Handle("/todos", isAuthorized(getTodos)).Methods("GET")
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
