// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package repository

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Item struct {
	ID          int32
	Site        string
	Price       string
	SmartTime   pgtype.Timestamp
	Name        string
	Description string
	Nickname    string
}