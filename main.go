package main

import (
	"github.com/joho/godotenv"
	"github.com/juseongkr/todo-list-go/server"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := server.ListenAndServe(server.Config{
		Address:     ":" + os.Getenv("PORT"),
		DatabaseURL: os.Getenv("DB_URL"),
	}); err != nil {
		log.Fatal(err)
	}
}
