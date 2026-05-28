package entity

import "time"

type Post struct {
	ID        string
	AuthorID  string
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
