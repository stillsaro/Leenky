// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package DBgo

import (
	"database/sql"
)

type User struct {
	ID           int32
	Username     string
	Email        string
	PasswordHash string
	CreatedAt    sql.NullTime
	UpdatedAt    sql.NullTime
	LastLogin    sql.NullTime
	Status       sql.NullString
}
