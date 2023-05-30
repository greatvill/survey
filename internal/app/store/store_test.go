package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

// calling one time
func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgres://user_db:123456@localhost:5437/restapi_test?sslmode=disable"
	}
	os.Exit(m.Run())
}
