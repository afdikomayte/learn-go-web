package domain

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

// type CustomerId uuid.UUID

type Customer struct {
	Id        uuid.UUID
	Name      string
	Email     sql.NullString
	Balance   int32
	Rating    float64
	BirthDate sql.NullTime
	Merried   bool
}

// interface customer usecase
type CustomerUseCase interface {
	Add(context context.Context, customer Customer) (*Customer, error)
}
