package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) route() http.Handler {
	g := gin.Default()

	v1 := g.Group("/api/v1")
	{
		v1.GET("/events", app.getAllEvent)
		v1.POST("/event", app.createEvent)
		v1.GET("/event/:id", app.getEvent)
		v1.PUT("/event/:id", app.updateEvent)
		v1.DELETE("/event/:id", app.deleteEvent)

		v1.POST("/event/:id/attendees/:userId", app.addAttendeeToEvent)
		v1.GET("/event/:id/attendees", app.getAttendeesForEvent)
		v1.DELETE("/event/:id/attendees/:userId", app.deleteAttendeesFromEvent)
		v1.GET("/attendee/:id/events", app.getEventByAttendees)

		v1.POST("/auth/register", app.registerUser)
	}

	return g
}
