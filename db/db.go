package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect(url string) error {
	c, err := sql.Open("postgres", url)
	if err != nil {
		return err
	}

	db = c
	return nil
}
