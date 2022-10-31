package models

import (
	"database/sql"
	"errors"
	"time"
)

type SnippetModelInterface interface {
	InsertSnippet(title string, content string, expires int) (int, error)
	GetSnippet(id int) (*Snippet, error)
	LatestSnippets() ([]*Snippet, error)
}

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// Define a SnippetModel type which wraps a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// This will insert a new snippet into the database.
func (m *SnippetModel) InsertSnippet(title, content string, expires int) (int, error) {
	query := `INSERT INTO snippets (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(query, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// This will return a specific snippet based on its id.
func (m *SnippetModel) GetSnippet(id int) (*Snippet, error) {
	query := `SELECT id, title, content, created, expires FROM snippets
    WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(query, id)

	// Initialze a pointer to a new zeroed Snippet struct
	snippet := &Snippet{}

	// Use row.Scan() to copy the values from each field in sql.Row to the
	// corresponding field in the Snippet struct.
	err := row.Scan(
		&snippet.ID,
		&snippet.Title,
		&snippet.Content,
		&snippet.Created,
		&snippet.Expires,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	return snippet, nil
}

// This will return the 10 most recently created snippets.
func (m *SnippetModel) LatestSnippets() ([]*Snippet, error) {
	query := `SELECT id, title, content, created, expires FROM snippets
    WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Initialze an empty slice to hold the Snippet structs.
	snippets := []*Snippet{}
	for rows.Next() {
		snippet := &Snippet{}
		err = rows.Scan(
			&snippet.ID,
			&snippet.Title,
			&snippet.Content,
			&snippet.Created,
			&snippet.Expires,
		)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, snippet)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
