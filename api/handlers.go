package api

import (
	"github.com/juseongkr/todo-list-go/db"
	"github.com/juseongkr/todo-list-go/todo"
	"net/http"
)

func getTodoLists(w http.ResponseWriter, r *http.Request) {
	lists, err := db.GetTodoLists()
	must(err)
	writeJSON(w, lists)
}

func createTodoList(w http.ResponseWriter, r *http.Request) {
	var req todo.List
	parseJSON(r.Body, &req)
	todoList, err := db.CreateTodoList(req.Name)
	must(err)
	writeJSON(w, todoList)
}

func getTodoList(w http.ResponseWriter, r *http.Request) {
	listID := parseIntParam(r, "list_id")
	list, err := db.GetTodoList(listID)
	must(err)
	writeJSON(w, list)
}

func modifyTodoList(w http.ResponseWriter, r *http.Request) {
	listID := parseIntParam(r, "list_id")
	var req todo.List
	parseJSON(r.Body, &req)
	must(db.RenameTodoList(listID, req.Name))
}

func deleteTodoList(w http.ResponseWriter, r *http.Request) {
	listID := parseIntParam(r, "list_id")
	must(db.DeleteTodoList(listID))
}

func createTodoItem(w http.ResponseWriter, r *http.Request) {
	listID := parseIntParam(r, "list_id")
	var req todo.Item
	parseJSON(r.Body, &req)

	item, err := db.CreateTodoItem(listID, req.Text, req.Done)
	must(err)
	writeJSON(w, item)
}

func modifyTodoItem(w http.ResponseWriter, r *http.Request) {
	listID := parseIntParam(r, "list_id")
	itemID := parseIntParam(r, "item_id")
	var req todo.Item
	parseJSON(r.Body, &req)
	must(db.ModifyTodoItem(listID, itemID, req.Text, req.Done))
}

func deleteTodoItem(w http.ResponseWriter, r *http.Request) {
	listID := parseIntParam(r, "list_id")
	itemID := parseIntParam(r, "item_id")
	must(db.DeleteTodoItem(listID, itemID))
}
