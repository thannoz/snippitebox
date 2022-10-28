package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

// Define a new UserModel type which wraps a database connection pool
type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) InsertUser(name, email, password string) error {
	return nil
}

// Authenticate verify whether a user with provided email & password exists
// it will return the relevant user ID if they do
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Exists checks if a user with a specific ID exists
func (m *UserModel) Exists(id int) (bool, error) {
	return false, nil
}
