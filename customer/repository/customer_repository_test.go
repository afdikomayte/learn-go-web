package repository_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/afdikomayte/learn-go-web/config"
	"github.com/afdikomayte/learn-go-web/customer/repository"
	"github.com/afdikomayte/learn-go-web/domain"
	"github.com/google/uuid"
)

var ctx = context.Background()

func TestInsertCustomer(t *testing.T) {
	customerRepo := repository.NewCustomerRepository(config.GetConnectionMysql(), ctx)

	//data for insert
	id, _ := uuid.NewUUID()
	//parse date
	layoutFormat := "2006-01-02 15:04:05"
	birthDay, _ := time.Parse(layoutFormat, "1995-09-02 00:00:00")

	customerData := domain.Customer{
		Id:        id,
		Name:      "afdiko 2",
		Email:     sql.NullString{String: "afdiko@bcid.com", Valid: true},
		Balance:   1000000000,
		Rating:    5.0,
		BirthDate: sql.NullTime{Time: birthDay, Valid: true},
		Merried:   true,
	}

	_, err := customerRepo.Insert(ctx, customerData)
	if err != nil {

		fmt.Println("Gagal Menyimpan data")

		panic(err)
	} else {
		fmt.Println("Berhasil Menyimpan data")

	}

}
