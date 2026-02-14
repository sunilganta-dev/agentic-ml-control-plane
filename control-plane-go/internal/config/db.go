package config

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func NewDB() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	var db *sql.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			time.Sleep(2 * time.Second)
			continue
		}

		err = db.Ping()
		if err == nil {
			return db, nil
		}

		fmt.Println("Waiting for database to be ready...")
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("database not ready after retries: %v", err)
}
