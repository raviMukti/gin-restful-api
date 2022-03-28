package app

import (
	"database/sql"
	"os"
	"time"

	"github.com/raviMukti/gin-restful-api/helper"
)

func NewDB() *sql.DB {
	connectionString := os.Getenv("DATABASE_USER") + ":" + os.Getenv("DATABASE_PASSWORD") + "@tcp(" + os.Getenv("DATABASE_URI") + ":" + os.Getenv("DATABASE_PORT") + ")/?parseTime=true"
	db, err := sql.Open("mysql", connectionString)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
