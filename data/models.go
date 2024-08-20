package data

import (
	"database/sql"
	"log"
	"os"
	"time"
)

var db *sql.DB

const dbTimeout = time.Second * 3

type Models struct {
	User User
}

func New(dbPool *sql.DB) Models {
	db = dbPool
	userModel := User{}

	CreateTable(
		"users",
	)

	return Models{
		User: userModel,
	}
}

// Open opens a connection to the database ans return the instance
func Open() *sql.DB {
	DB_URI := os.Getenv("DNS")
	db, err := sql.Open("postgres", DB_URI)
	if err != nil {
		log.Panicln("Error connecting to database")
		return nil
	}

	return db
}
