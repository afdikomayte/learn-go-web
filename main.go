package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/afdikomayte/learn-go-web/customer/controller/web/customer"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		type DataPage map[string]interface{}
		dataPage := DataPage{"pageName": "Home"}
		tmpl := template.Must(template.ParseFiles(
			"customer/views/index/index.html",
			"customer/views/partial/_header.html",
			"customer/views/partial/_navbar.html",
			"customer/views/partial/_footer.html",
		))

		err := tmpl.ExecuteTemplate(w, "index", dataPage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("/customer/add", customer.HandelCustomAdd)
	mux.HandleFunc("/customer/store", customer.HandelCustomerStore)

	fmt.Println("server run at http://localhost:8080")

	//server
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
