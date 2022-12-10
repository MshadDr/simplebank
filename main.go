package main

import (
	"database/sql"
	"log"
	"simplebank/api"
	database "simplebank/database/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:root@localhost:5432/simple_database?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	store := database.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)

}