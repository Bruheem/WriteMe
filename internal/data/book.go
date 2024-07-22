package data

import (
	"database/sql"
	"errors"
	"time"
)

type Chapter struct {
	ID      int
	Title   string
	Content []byte
}

type Book struct {
	ID         int
	Title      string
	Author     string
	Created_at time.Time
	Chapters   []Chapter
}

type BookModel struct {
	DB *sql.DB
}

func (m BookModel) Get(id int64) (*Book, error) {
	return nil, nil
}

func (m BookModel) Insert(b *Book) error {

	query := `INSERT INTO book (title, author) VALUES ($1, $2) RETURNING id`
	args := []interface{}{b.Title, b.Author}

	return m.DB.QueryRow(query, args...).Scan(&b.ID)
}

func (m BookModel) Update(b *Book) error {
	return nil
}

func (m BookModel) Delete(id int64) error {

	query := `DELETE FROM book WHERE id = $1`

	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows were affected in the database!")
	}

	return nil
}
