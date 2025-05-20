package database

import (
	"database/sql"

	"github.com/rajan-marasini/event-api/internal/models"
)

type Models struct {
	User     models.UserModel
	Event    models.EventModel
	Attendee models.AttendeeModel
}

func NewModel(db *sql.DB) Models {
	return Models{
		User:     models.UserModel{DB: db},
		Event:    models.EventModel{DB: db},
		Attendee: models.AttendeeModel{DB: db},
	}
}
