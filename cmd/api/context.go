package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rajan-marasini/event-api/internal/models"
)

func (app *application) GetUserFromContext(c *gin.Context) *models.User {
	contextUser, exist := c.Get("user")
	if !exist {
		return &models.User{}
	}

	user, ok := contextUser.(*models.User)
	if !ok {
		return &models.User{}
	}
	return user
}
