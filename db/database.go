package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Env holds database connection to Postgres
type Env struct {
	DB *sql.DB
}


// database variables
// use dotenv to get these
const (
	host = "localhost"
	port = 5439
	user = "clama"
	password = "clama17565"
	dbname = "artists"
)

// Connect DB tries to connect DB and on success it returns
// DB connection string and nil error, otherwise it returns empty DB and the corresponding error.
func ConnectDB() (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Printf("failed to connect to database: %v", err)
		return &sql.DB{}, err
	}
	return db, nil
}

