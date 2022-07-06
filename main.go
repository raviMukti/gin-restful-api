package main

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/raviMukti/gin-restful-api/app"
	"github.com/raviMukti/gin-restful-api/messaging"
)

func main() {

	app.Init()
	messaging.InitKafkaConfig()

	router := app.SetupRouter()

	router.Run(":" + os.Getenv("APP_PORT"))
}
