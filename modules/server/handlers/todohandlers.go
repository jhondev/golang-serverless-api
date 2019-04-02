package handlers

import (
	"net/http"
	"serverlessapi/modules/todo"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func GetATodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	todo := todo.GetAToDo(todoID)
	render.JSON(w, r, todo) // A chi router helper for serializing and returning json
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	response := todo.DeleteTodo("todoID")
	render.JSON(w, r, response) // Return some demo response
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	response := todo.CreateTodo()
	render.JSON(w, r, response) // Return some demo response
}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos := todo.GetAllTodos()
	render.JSON(w, r, todos) // A chi router helper for serializing and returning json
}
