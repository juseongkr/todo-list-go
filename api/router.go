package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func TodoListAPI() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/lists", getTodoList).Methods(http.MethodGet)

	router.HandleFunc("/list", createTodoList).Methods(http.MethodPost)
	router.HandleFunc("/list/{list_id}", getTodoList).Methods(http.MethodGet)
	router.HandleFunc("/list/{list_id}", modifyTodoList).Methods(http.MethodPut)
	router.HandleFunc("/list/{list_id}", deleteTodoList).Methods(http.MethodDelete)

	router.HandleFunc("/list/{list_id}/item", createTodoItem).Methods(http.MethodPost)
	router.HandleFunc("/list/{list_id}/item/{item_id}", modifyTodoItem).Methods(http.MethodPut)
	router.HandleFunc("/list/{list_id}/item/{item_id}", deleteTodoItem).Methods(http.MethodDelete)

	return router
}
