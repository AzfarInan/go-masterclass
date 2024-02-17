package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/AzfarInan/go-masterclass/simplebank/db/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config: " + err.Error())
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		panic("cannot connect to db: " + err.Error())
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
