package data

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type Dal struct {
	DB *sql.DB
}

func ConnectSQLDb() *sql.DB {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("libsql", url)
	if err != nil {
		log.Fatalf("Failed to open db, error: %s", err)
		os.Exit(1)
	}
	return db
}
