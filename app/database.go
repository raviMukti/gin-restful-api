package app

import (
	"database/sql"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/raviMukti/gin-restful-api/helper"
)

func InitDatabase() *sql.DB {

	err := godotenv.Load()
	helper.PanicIfError(err)

	connectionString := os.Getenv("DATABASE_USER") + ":" + os.Getenv("DATABASE_PASSWORD") + "@tcp(" + os.Getenv("DATABASE_URI") + ":" + os.Getenv("DATABASE_PORT") + ")/" + os.Getenv("DATABASE_INSTANCE") + "?parseTime=true"
	db, err := sql.Open("mysql", connectionString)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
