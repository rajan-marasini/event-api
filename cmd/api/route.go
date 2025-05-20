package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) route() http.Handler {
	g := gin.Default()

	return g
}
