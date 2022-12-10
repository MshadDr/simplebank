package main

import (
	"database/sql"
	"log"
	"simplebank/api"
	database "simplebank/database/sqlc"
	"simplebank/utils"
	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cann't load config...")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	store := database.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

}