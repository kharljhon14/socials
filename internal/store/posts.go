package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Post struct {
	ID        uuid.UUID `json:"id"`
	Content   string    `json:"content"`
	Title     string    `json:"title"`
	UserID    uuid.UUID `json:"user_id"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostgresPostsStore struct {
	db *sql.DB
}

func (s *PostgresPostsStore) Create(ctx context.Context, post *Post) error {
	query := `
		INSERT INTO posts
		(content, title, user_id, tags)
		VALUES
		($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`

	args := []any{
		post.Content,
		post.Title,
		post.UserID,
		pq.Array(post.Tags),
	}

	err := s.db.QueryRowContext(ctx, query, args...).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
