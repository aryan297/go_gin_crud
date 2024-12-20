package models

import (
	"time"

	"event.com/first/db"
)

type Event struct {
	Id          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserId      int
}

//var events = []Event{}

func (e *Event) StoreEvents() error {
	query := `INSERT INTO events("name","description","location","dateTime","user_id") VALUES(?,?,?,?,?)`

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
	if err != nil {
		return err
	}
	e.Id = id
	return nil
}

func GetEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)

	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	/* 	for i := 0; i < len(events); i++ {
		if events[i].Id == id {
			return events[i]
		}
	} */
	query := `SELECT * FROM events WHERE id = ?`
	rows := db.DB.QueryRow(query, id)

	var event Event
	rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

	return &event, nil
}

func (event Event) UpdateEvents() (*Event, error) {
	query := `UPDATE events SET name=?,description=?,location=?,dateTime=? WHERE id=?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.Id)
	return &event, err

}

func (event Event) DeleteEvents() error {
	query := `DELETE FROM events WHERE id=?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.Id)
	return err
}
