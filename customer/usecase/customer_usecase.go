package usecase

import (
	"context"

	"github.com/afdikomayte/learn-go-web/customer/repository"
	"github.com/afdikomayte/learn-go-web/domain"
)

type customerUseCase struct {
	customerRepo repository.CustomerRepository
}

// func yang akan di gunakan di setiap layer
func NewCustomerUseCase(customerRepo repository.CustomerRepository) domain.CustomerUseCase {
	return &customerUseCase{
		customerRepo: customerRepo,
	}
}

func (cu *customerUseCase) Add(ctx context.Context, customer domain.Customer) (res *domain.Customer, err error) {
	res, err = cu.Add(ctx, customer)

	return res, err
}
