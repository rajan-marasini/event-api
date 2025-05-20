package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type EventModel struct {
	DB *sql.DB
}

type Event struct {
	ID          int    `json:"id"`
	OwnerId     int    `json:"owner_id"`
	Name        string `json:"name" binding:"required,min=3"`
	Description string `json:"description" binding:"required,min=10"`
	Date        string `json:"date" binding:"required"`
	Location    string `json:"location" binding:"required,min=3"`
}

func (m *EventModel) Insert(event *Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	fmt.Println(&event)

	query := `
		INSERT INTO events (owner_id, name, description, location, date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	return m.DB.QueryRowContext(ctx, query, event.OwnerId, event.Name, event.Description, event.Location, event.Date).Scan(&event.ID)
}

func (m *EventModel) Get(id int) (*Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var event Event

	query := `
		SELECT 
		id, owner_id, name, description, location, date
		FROM events WHERE id=$1 LIMIT 1
	`

	err := m.DB.QueryRowContext(ctx, query, id).Scan(&event.ID, &event.OwnerId, &event.Name, &event.Description, &event.Location, &event.Date)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &event, nil
}

func (m *EventModel) GetAll() ([]*Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var events []*Event

	query := "SELECT * FROM events"

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var event Event

		err := rows.Scan(&event.ID, &event.OwnerId, &event.Name, &event.Description, &event.Location, &event.Date)
		if err != nil {
			return nil, err
		}
		events = append(events, &event)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}

func (m *EventModel) Update(event *Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		UPDATE events
		SET name=$1, description=$2, date=$3, location=$4
		WHERE id=$5
	`
	_, err := m.DB.ExecContext(ctx, query, event.Name, event.Description, event.Date, event.Location, event.ID)
	if err != nil {
		return err
	}
	return nil
}

func (m *EventModel) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		DELETE FROM events WHERE id=$1
	`
	_, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
