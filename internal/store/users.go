package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"_"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"upated_at"`
}

type PostgresUsersStore struct {
	db *sql.DB
}

func (s *PostgresUsersStore) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users
		(username, email, password)
		VALUES
		($1, $2, $3)
		RETURNING id, created_at, updated_at
	`

	args := []any{
		user.Username,
		user.Email,
		user.Password,
	}

	err := s.db.QueryRowContext(ctx, query, args...).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
