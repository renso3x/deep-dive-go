// REST API using mux

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Todo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var todos []Todo

func allTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// get the body of our POST request
	// unmarshal this into a new Todo struct
	// append this to our Todos array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	// _ = json.NewDecoder(r.Body).Decode(&todo)

	var todo Todo

	todo.Id = strconv.Itoa(rand.Intn(10000))
	json.Unmarshal(reqBody, &todo)

	todos = append(todos, todo)

	json.NewEncoder(w).Encode(&todo)
}

func fetchTodoById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]

	for _, t := range todos {
		if t.Id == key {
			json.NewEncoder(w).Encode(t)
		}
	}
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)

	for index, item := range todos {
		if item.Id == id {
			// delete first the index
			todos = append(todos[:index], todos[index+1:]...)
			// update the todo
			var todo Todo
			json.Unmarshal(reqBody, &todo)
			todo.Id = id
			todos = append(todos, todo)
			break

		}
	}
	json.NewEncoder(w).Encode(todos)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	// GET THE PARAMS
	params := mux.Vars(r)
	id := params["id"]

	for index, t := range todos {
		if t.Id == id {
			// delete the todo
			todos = append(todos[:index], todos[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(todos)
}

func handleRequest() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/todos", allTodos).Methods("GET")
	r.HandleFunc("/todos", addTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", fetchTodoById).Methods("GET")
	r.HandleFunc("/todos/{id}", updateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", r))
}

func main() {
	todos = []Todo{
		Todo{Id: "1", Name: "Watch NBA"},
		Todo{Id: "2", Name: "Drink Coffe"},
		Todo{Id: "3", Name: "Wash clothes"},
	}
	handleRequest()
}
