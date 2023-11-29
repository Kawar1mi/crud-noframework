package db

import (
	"database/sql"
	"os"

	loghelper "github.com/Kawar1mi/crud-noframework/internal/log_helper"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func DatabaseConnection() *sql.DB {

	// DSN in format - "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
	db, err := sql.Open("postgres", os.Getenv("DSN"))
	if err != nil {
		loghelper.FatalIfError(err)
	}

	err = db.Ping()
	if err != nil {
		loghelper.FatalIfError(err)
	}

	loghelper.InfoMsg("connected to database")

	return db
}
