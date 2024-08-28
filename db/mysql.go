package db

import (
	"database/sql"
	"log"

	"github.com/kodinggo/user-service-gb1/internal/helper"

	_ "github.com/go-sql-driver/mysql"
)

// NewMysql is a function to initialize the MySQL database.
func NewMysql() *sql.DB {
	// Initialize the database
	db, err := sql.Open("mysql", helper.GetConnectionString())
	if err != nil {
		log.Fatal(err)
	}

	// Check database connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
