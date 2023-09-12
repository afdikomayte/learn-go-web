package customer

import (
	"database/sql"
	"errors"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/afdikomayte/learn-go-web/config"
	"github.com/afdikomayte/learn-go-web/customer/repository"
	"github.com/afdikomayte/learn-go-web/domain"
	"github.com/google/uuid"
)

type DataPage map[string]interface{}

func HandelCustomAdd(w http.ResponseWriter, r *http.Request) {
	//data page
	dataPage := DataPage{"pageName": "Add Customer"}

	var tmpl = template.Must(template.ParseFiles(
		"customer/views/add.html",
		"customer/views/partial/_header.html",
		"customer/views/partial/_footer.html",
		"customer/views/partial/_navbar.html",
	))

	err := tmpl.ExecuteTemplate(w, "add", dataPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HandelCustomerStore(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	customerRepo := repository.NewCustomerRepository(config.GetConnectionMysql(), ctx)

	if r.Method == "POST" {
		// convers data dari form sesuai data struct
		email := r.FormValue("email")
		balance, err := strconv.Atoi(r.FormValue("balance"))
		if err != nil {
			panic(errors.New("balance " + err.Error()))
		}
		rating, err := strconv.ParseFloat(r.FormValue("rating"), 64)
		if err != nil {
			panic(errors.New("rating " + err.Error()))
		}
		layoutFormat := "2006-01-02 15:04:05"

		birthDateForm := r.FormValue("birth_date") + " 00:00:00"
		birthDate, err := time.Parse(layoutFormat, birthDateForm)
		if err != nil {
			panic(errors.New("birthDay" + err.Error()))
		}
		isMarried := r.Form["isMarried"]

		var merried bool
		if len(isMarried) == 0 {
			merried = false
		} else if isMarried[0] != "1" {
			merried = false
		} else {
			merried = true
		}

		//passing data form to struct
		customer := domain.Customer{
			Id:        uuid.New(),
			Name:      r.FormValue("name"),
			Email:     sql.NullString{String: email, Valid: true},
			Balance:   int32(balance),
			Rating:    float64(rating),
			BirthDate: sql.NullTime{Time: birthDate, Valid: true},
			Merried:   merried,
		}

		//insert to database
		_, errs := customerRepo.Insert(ctx, customer)

		if errs != nil {

			http.Error(w, err.Error(), http.StatusInternalServerError)

			panic(errs)
		}
		http.Redirect(w, r, "http://localhost:8080/add", http.StatusSeeOther)
	}
	http.Error(w, "", http.StatusInternalServerError)
}
