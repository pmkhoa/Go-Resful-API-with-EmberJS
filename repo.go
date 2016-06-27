package main

import (
	"fmt"
	"strconv"
)

var currentId int
var todos Todos

func init() {
	RepoCreateTodo(Todo{Name: "Test Todo 1"})
	RepoCreateTodo(Todo{Name: "Test Todo 2"})
}

func RepoCreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = strconv.Itoa(currentId)
	todos = append(todos, t)
	return t
}

func RepoFindTodo(id string) Todo {
	for _, todo := range todos {
		if todo.Id == id {
			return todo
		}
	}
	return Todo{}
}

func RepoDeleteTodo(id string) error {
	for i, todo := range todos {
		if todo.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}

func RepoUpdateTodo(id string, newTodoData Todo) error {
	for index, todo := range todos {
		if todo.Id == id {
			updateTodoData(&todos[index], newTodoData)
			return nil
		}
	}
	return fmt.Errorf("Cannot update Todo")
}

func updateTodoData(srcTodo *Todo, newData Todo) error {
	if newData.Name != "" {
		srcTodo.Name = newData.Name
	}
	if srcTodo.Completed != newData.Completed {
		srcTodo.Completed = newData.Completed
	}
	if srcTodo.Due != newData.Due {
		srcTodo.Due = newData.Due
	}
	return nil
}
