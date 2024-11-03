// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package DBgo

import (
	"context"
)

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, username, email, password_hash, created_at, updated_at, last_login, status 
FROM users 
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastLogin,
		&i.Status,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, username, email, password_hash, created_at, updated_at, last_login, status 
FROM users 
WHERE username = $1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastLogin,
		&i.Status,
	)
	return i, err
}

const insertUser = `-- name: InsertUser :exec
INSERT INTO users (
    username, email, password_hash
) VALUES ($1, $2, $3)
`

type InsertUserParams struct {
	Username     string
	Email        string
	PasswordHash string
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) error {
	_, err := q.db.ExecContext(ctx, insertUser, arg.Username, arg.Email, arg.PasswordHash)
	return err
}
