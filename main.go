package main

import (
	"database/sql"
	"github.com/Geun-Oh/backend/api"
	db "github.com/Geun-Oh/backend/db/sqlc"
	"github.com/joho/godotenv"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbSource := os.Getenv("DB_SOURCE")

	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
}
