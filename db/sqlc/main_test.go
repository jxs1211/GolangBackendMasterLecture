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

func TestMain(m *testing.M) {

	conn, err := sql.Open(driver, openSource)
	if err != nil {
		log.Fatal(err)
		return
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}
