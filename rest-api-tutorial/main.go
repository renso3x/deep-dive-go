package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	// title       string `json:"title"`
	// description string `json:"description"`
	// author      string `json:"author"`

	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func fetchArticles(w http.ResponseWriter, r *http.Request) {

	// articles := []Article{
	// 	Article{title: "Full Stack Development", description: "This is a sample decription", author: "Romeo Enso"},
	// 	Article{title: "Go RestAPI", description: "This is a sample decription", author: "Romeo Enso"},
	// }

	articles := []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my first go")
}

func postArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test POst endpoint worked")
}

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/articles", fetchArticles).Methods("GET")
	router.HandleFunc("/articles", postArticles).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	handleRequest()
}
