package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/nvmh0103/bank-golang/api"
	db "github.com/nvmh0103/bank-golang/db/sqlc"
	"github.com/nvmh0103/bank-golang/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load configuration", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
}
