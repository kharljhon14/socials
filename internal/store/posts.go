package store

import (
	"context"
	"database/sql"
	"errors"
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

func (s PostgresPostsStore) GetByID(ctx context.Context, postID uuid.UUID) (*Post, error) {
	query := `
		SELECT 
			id,
			title,
			content,
			userID,
			tags,
			created_at,
			updated_at
		FROM posts
		WHERE id = $1
	`

	var post Post
	err := s.db.QueryRowContext(ctx, query, postID).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.UserID,
		pq.Array(&post.Tags),
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &post, nil
}
