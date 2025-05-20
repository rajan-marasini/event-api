package models

import "database/sql"

type AttendeeModel struct {
	DB *sql.DB
}

type Attendee struct {
	ID      int `json:"id"`
	UserId  int `json:"user_id"`
	EventId int `json:"event_id"`
}
