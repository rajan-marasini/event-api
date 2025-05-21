package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (app *application) route() http.Handler {
	g := gin.Default()

	v1 := g.Group("/api/v1")
	{
		v1.GET("/events", app.getAllEvent)
		v1.GET("/event/:id", app.getEvent)

		v1.GET("/event/:id/attendees", app.getAttendeesForEvent)
		v1.GET("/attendee/:id/events", app.getEventByAttendees)

		v1.POST("/auth/register", app.registerUser)
		v1.POST("/auth/login", app.login)
	}

	authGroup := v1.Group("/")
	authGroup.Use(app.AuthMiddleeware())
	{
		authGroup.POST("/event", app.createEvent)
		authGroup.PUT("/event/:id", app.updateEvent)
		authGroup.DELETE("/event/:id", app.deleteEvent)

		authGroup.POST("/event/:id/attendees/:userId", app.addAttendeeToEvent)
		authGroup.DELETE("/event/:id/attendees/:userId", app.deleteAttendeesFromEvent)
	}

	g.GET("/swagger/*any", func(c *gin.Context) {
		if c.Request.RequestURI == "/swagger/" {
			c.Redirect(302, "/swagger/index.html")
		}
		ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("http://localhost:8000/swagger/doc.json"))(c)
	})

	return g
}
