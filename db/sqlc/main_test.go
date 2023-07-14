package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	driver     = "postgres"
	openSource = "postgres://root:root@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(driver, openSource)
	if err != nil {
		log.Fatal(err)
		return
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
