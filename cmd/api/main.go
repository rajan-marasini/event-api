package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rajan-marasini/event-api/internal/database"
)

type application struct {
	port      int
	jwtSecret string
	models    database.Models
}

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}

	models := database.NewModel(db)

	app := &application{
		port:      8000,
		jwtSecret: "jwt_secret_key",
		models:    models,
	}

	if err := app.serve(); err != nil {
		panic(err)
	}
}
