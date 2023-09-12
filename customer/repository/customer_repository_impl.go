package repository

import (
	"context"
	"database/sql"

	"github.com/afdikomayte/learn-go-web/domain"
)

type CustomerRepositoryImpl struct {
	DB  *sql.DB
	ctx context.Context
}

func NewCustomerRepository(db *sql.DB, ctx context.Context) CustomerRepository {
	return &CustomerRepositoryImpl{DB: db, ctx: ctx}
}

func (c *CustomerRepositoryImpl) Insert(ctx context.Context, customer domain.Customer) (*domain.Customer, error) {

	//sql script
	scriptSql := `INSERT INTO customer_detail(id,name,email,balance,rating,birth_date,merried) VALUES (?,?,?,?,?,?,?)`

	_,err := c.DB.ExecContext(ctx,scriptSql,customer.Id,customer.Name,customer.Email,customer.Balance,customer.Rating,customer.BirthDate,customer.Merried)
	if err != nil{
		return &domain.Customer{},err
	}
	
	return &domain.Customer{},nil
}
