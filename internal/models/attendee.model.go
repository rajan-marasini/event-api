package models

import (
	"context"
	"database/sql"
	"time"
)

type AttendeeModel struct {
	DB *sql.DB
}

type Attendee struct {
	ID      int `json:"id"`
	UserId  int `json:"user_id"`
	EventId int `json:"event_id"`
}

func (m *AttendeeModel) Insert(attendee *Attendee) (*Attendee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		INSERT INTO attendees (event_id, user_id)
		VALUES ($1, $2)
		RETURNING id
	`
	err := m.DB.QueryRowContext(ctx, query, attendee.EventId, attendee.UserId).Scan(&attendee.ID)
	if err != nil {
		return nil, err
	}

	return attendee, nil
}

func (m *AttendeeModel) GetByEventAndAttendee(eventId, userId int) (*Attendee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		SELECT id, event_id, user_id
		FROM attendees
		WHERE event_id = $1 AND user_id = $2
	`
	var attendee Attendee
	err := m.DB.QueryRowContext(ctx, query, eventId, userId).Scan(&attendee.ID, &attendee.EventId, &attendee.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &attendee, nil
}

func (m *AttendeeModel) GetAttendeesByEvent(eventId int) ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		SELECT u.id, u.name, u.email
		FROM users u
		JOIN attendees a ON u.id = a.user_id
		WHERE a.event_id = $1
	`
	rows, err := m.DB.QueryContext(ctx, query, eventId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}
