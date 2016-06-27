package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/todos", TodosIndex)
	router.GET("/todos/:todoId", ShowTodo)
	router.DELETE("/todos/:todoId", DeleteTodo)
	router.PUT("/todos/:todoId", UpdateTodo)
	router.PATCH("/todos/:todoId", UpdateTodo)
	router.POST("/todos", CreateTodo)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "PATCH", "OPTIONS"},
	})

	log.Fatal(http.ListenAndServe(":8080", c.Handler(router)))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello World")
}

func TodosIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(todos)
	// fmt.Printf("%+v\n", json)
	if err != nil {
		panic(err)
	}
}

func ShowTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todoId := params.ByName("todoId")
	todo := RepoFindTodo(todoId)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(todo)
	if err != nil {
		panic(err)
	}
}

// curl -H "Content-Type: application/json" -d '{"name":"New Todo"}'
// http://localhost:8080"
func CreateTodo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var todo Todo
	//Read Inputs
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &todo)
	if err != nil {
		panic(err)
	}
	t := RepoCreateTodo(todo)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(t)
	if err != nil {
		panic(err)
	}
}

// curl -X DELETE 'http://localhost:8080/todos/2'
func DeleteTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todoId := params.ByName("todoId")
	err := RepoDeleteTodo(todoId)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(nil)
	if err != nil {
		panic(err)
	}
}

// curl -H "Accept: application/json" -H "Content-type: application/json"
// -X PUT -d '{"name": "Todo 2 Updated 2"}' 'http://localhost:8080/todos/2'
func UpdateTodo(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todoId := params.ByName("todoId")
	var newTodoData Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &newTodoData)
	if err != nil {
		panic(err)
	}

	err = RepoUpdateTodo(todoId, newTodoData)

	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(nil)
}
