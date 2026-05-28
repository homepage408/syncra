package sqlc

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Queries struct {
	db *sql.DB
}

type User struct {
	ID           string
	Email        string
	PasswordHash string
	Role         string
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Post struct {
	ID        string
	AuthorID  string
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateUserParams struct {
	ID           string
	Email        string
	PasswordHash string
	Role         string
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type CreatePostParams struct {
	ID        string
	AuthorID  string
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(db *sql.DB) *Queries {
	return &Queries{db: db}
}

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	return nil, errors.New("not implemented")
}

func (q *Queries) CreateUser(ctx context.Context, params CreateUserParams) error {
	return errors.New("not implemented")
}

func (q *Queries) GetPostByID(ctx context.Context, id string) (*Post, error) {
	return nil, errors.New("not implemented")
}

func (q *Queries) ListPosts(ctx context.Context) ([]*Post, error) {
	return nil, errors.New("not implemented")
}

func (q *Queries) CreatePost(ctx context.Context, params CreatePostParams) error {
	return errors.New("not implemented")
}
