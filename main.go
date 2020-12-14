package main

import (
	"github.com/juseongkr/todo-list-go/server"
	"github.com/joho/godotenv"
	"log"
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
