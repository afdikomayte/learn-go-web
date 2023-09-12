package repository

import (
	"context"

	"github.com/afdikomayte/learn-go-web/domain"
)

type CustomerRepository interface {
	Insert(ctx context.Context, customer domain.Customer) (*domain.Customer, error)
}
