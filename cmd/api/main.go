package main

import (
	"database/sql"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rajan-marasini/event-api/internal/database"
	"github.com/rajan-marasini/event-api/internal/env"
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
		port:      env.GetEnvInt("PORT", 8000),
		jwtSecret: env.GetEnvString("JWT_SECRET", "kjhfajkhjakfhjkadf"),
		models:    models,
	}

	if err := app.serve(); err != nil {
		panic(err)
	}
}
