package config_test

import (
	"fmt"
	"testing"

	getConnection "github.com/afdikomayte/learn-go-web/config"
)

func TestDatabaseConnetion(t *testing.T) {
	db := getConnection.GetConnectionMysql()
	defer db.Close()

	fmt.Println("database connection success")
}

func GetConnetionMysql() {
	panic("unimplemented")
}
