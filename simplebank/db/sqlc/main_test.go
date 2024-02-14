package db

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgresql://:@localhost:5432/simplebank?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error

	testDB, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		panic("cannot connect to db: " + err.Error())
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
