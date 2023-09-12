package customer_test

import (
	"fmt"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/afdikomayte/learn-go-web/customer/controller/web/customer"
)

func TestAddCustomer(t *testing.T) {
	requestBody := strings.NewReader("name=testhandel&email=afdiko@bcid.com&balance=10000&rating=5.0&birth_date=1995-09-02 00:00:00&isMarried=true")
	request := httptest.NewRequest("POST", "http://localhost/add", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	customer.HandelCustomerStore(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
