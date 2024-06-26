package models

import (
	"project/REST_API/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64
}

type RegisteredEvent struct {
	ID      int64
	EventId int64
	UserId  int64
}

func (e *Event) Save() error {

	query := `INSERT INTO events (name, description,location,dateTime,user_id)
			  VALUES(?,?,?,?,?) 		
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id
	return err

}

func (event Event) DeleteEvent() error {

	query := `DELETE FROM events WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID)

	return err
}

func GetAllEvents() ([]Event, error) {

	query := `SELECT * FROM events`

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil

}

func GetAllRegisteredEvents() ([]RegisteredEvent, error) {
	query := `SELECT * FROM registerations`

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var registeredEvents []RegisteredEvent
	for rows.Next() {
		var registeredEvent RegisteredEvent
		err := rows.Scan(&registeredEvent.ID, &registeredEvent.EventId, &registeredEvent.UserId)
		if err != nil {
			return nil, err
		}
		registeredEvents = append(registeredEvents, registeredEvent)
	}

	return registeredEvents, nil
}

func GetEventById(id int64) (*Event, error) {

	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}
	return &event, nil

}

func (event Event) Update() error {

	query := `
	UPDATE events
	SET name = ? , description = ? , location = ? , datetime = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err

}

func New(name string, desc string, loc string, dt time.Time, userId int64) Event {
	return Event{
		Name:        name,
		Description: desc,
		Location:    loc,
		DateTime:    dt,
		UserId:      userId,
	}
}

func (e Event) Register(userId int64) error {

	query := "INSERT INTO registerations (event_id, user_id) VALUES(?,?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)
	if err != nil {
		return err
	}

	return nil
}

func (e Event) CancelRegisteration(userId int64) error {
	query := `DELETE FROM registerations WHERE event_id = ? AND user_id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}
