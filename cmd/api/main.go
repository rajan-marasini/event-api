package main

import (
	"database/sql"

	_ "github.com/rajan-marasini/event-api/docs"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rajan-marasini/event-api/internal/database"
	"github.com/rajan-marasini/event-api/internal/env"
)

// @title Go Gin Rest API
// @version 1.0.0
// @description A rest API in Go using Gin framework
// @securityDefinition.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter your bearer token in the format **Bearer &lt;token&gt;**

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
