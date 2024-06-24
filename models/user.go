package models

import (
	"fmt"
	"project/REST_API/db"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error {

	query := "INSERT INTO users (email,password) VALUES (?,?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare query: %w", err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to retrieve the last user id: %w ", err)
	}

	user.ID = userID
	return err

}
