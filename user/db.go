package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		dbURL = "postgresql://postgres@localhost:5432/postgres?sslmode=disable"
	}
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("user db connection err: %v", err)
	}

	if err := conn.Ping(); err != nil {
		log.Fatalf("failed to ping user db err: %v", err)
	}

	db = conn
}
