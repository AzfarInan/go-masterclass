package main

import (
	"database/sql"
	"log"

	"github.com/AzfarInan/go-masterclass/simplebank/api"
	db "github.com/AzfarInan/go-masterclass/simplebank/db/sqlc"
	"github.com/AzfarInan/go-masterclass/simplebank/db/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: " + err.Error())
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: " + err.Error())
	}
}
