package main

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/raviMukti/gin-restful-api/app"
)

func main() {

	app.Init()

	router := app.SetupRouter()

	router.Run(":" + os.Getenv("APP_PORT"))
}
