package main

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
)

const (
	MIGRATION_DIR = "migrate/migrations"
)

func main() {
	if len(os.Args) < 2 {
		panic("Please provide migration direction 'up' or 'down'")
	}

	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}

	direction := os.Args[1]

	if err := goose.SetDialect("sqlite3"); err != nil {
		panic(err)
	}

	switch direction {
	case "up":
		if err := goose.Up(db, MIGRATION_DIR); err != nil {
			panic(err)
		}
	case "down":
		if err := goose.Down(db, MIGRATION_DIR); err != nil {
			panic(err)
		}
	default:
		panic("Invalid direction. Use 'up' or 'down'")
	}

}
