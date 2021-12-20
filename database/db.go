package database

import (
	"fmt"
	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@postgres-postgresql:5432/postgres?sslmode=disable"
)

//When using dockers local host of each docker is the container name!

func NewDB() (*sqlx.DB, error) {
	dbSource := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_SSL_MODE"),
	)

	db, err := sqlx.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
		return db, err
	}

	return db, nil
}
