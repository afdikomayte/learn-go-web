package handler

import (
	"fmt"
	"net/http"
)

func GetCustomer (w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "")
}
