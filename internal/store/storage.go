package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

var ErrNotFound = errors.New("record not found")

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetByID(context.Context, uuid.UUID) (*Post, error)
	}
	Comments interface {
		GetByPostID(context.Context, uuid.UUID) ([]Comment, error)
	}
	Users interface {
		Create(context.Context, *User) error
	}
}

func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		Posts:    &PostgresPostsStore{db: db},
		Comments: &CommentsStore{db: db},
		Users:    &PostgresUsersStore{db: db},
	}
}
