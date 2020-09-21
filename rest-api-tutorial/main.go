package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	title       string   "json:title"
	description string   "json:description"
	author      string   "json:author"
	tags        []string "json:tags"
}

type Articles []Article

func fetchArticles(w http.ResponseWriter, _ *http.Request) {
	tags := []string{"golang", "programming"}

	articles := Articles{
		Article{title: "Full Stack Development", description: "This is a sample decription", author: "Romeo Enso", tags: tags},
		Article{title: "Go RestAPI", description: "This is a sample decription", author: "Romeo Enso", tags: tags},
	}

	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("Welcome to my first go")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", fetchArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequest()
}
