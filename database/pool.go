package database

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"

	"log"
)

// ToDo : Get config from yaml
type Config struct {
	Host            string // DB URI
	Port            int
	User            string
	Password        string
	DBName          string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

func CreateDBConnectionPool() *sql.DB {
	connStr := "host=localhost port=5432 user=test password=AdminSex dbname=mydatabase sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	// Test connection
	if err := db.Ping(); err != nil {
		log.Fatal("Cannot ping database:", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	return db
}
