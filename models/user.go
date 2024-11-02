package models

import "event.com/first/db"

type User struct {
	id       int64
	email    string `binding:required`
	password string `binding:required`
}

func (u User) Save() error {
	query := `INSERT INTO users(email, password) VALUES (?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(u.email, u.password)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.id = userId
	return nil
}
