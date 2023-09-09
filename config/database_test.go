package config

import (
	"fmt"
	"testing"
)

func TestDatabaseConnetion(t *testing.T) {
	db := GetConnetionMysql()
	defer db.Close()

	fmt.Println("database connection success")
}
