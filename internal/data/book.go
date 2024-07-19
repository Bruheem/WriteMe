package data

import "time"

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
