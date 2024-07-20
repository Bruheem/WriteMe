package data

import (
	"database/sql"
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
	return nil
}

func (m BookModel) Update(b *Book) error {
	return nil
}

func (m BookModel) Delete(b *Book) error {
	return nil
}
