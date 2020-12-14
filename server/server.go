package server

import (
	"github.com/juseongkr/todo-list-go/api"
	"github.com/juseongkr/todo-list-go/db"
	"net/http"
)

type Config struct {
	Address     string
	DatabaseURL string
}

func ListenAndServe(cfg Config) error {
	if err := db.Connect(cfg.DatabaseURL); err != nil {
		return err
	}

	return http.ListenAndServe(cfg.Address,
		loggingMiddleware(api.TodoListAPI()))
}
