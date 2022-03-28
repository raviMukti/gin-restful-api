package main

import (
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/raviMukti/gin-restful-api/app"
	"github.com/raviMukti/gin-restful-api/controller"
	"github.com/raviMukti/gin-restful-api/helper"
	"github.com/raviMukti/gin-restful-api/repository"
	"github.com/raviMukti/gin-restful-api/service"
)

func main() {

	err := godotenv.Load()
	helper.PanicIfError(err)

	db := app.NewDB()
	validate := validator.New()
	carRepository := repository.NewCarRepository()
	carService := service.NewCarService(carRepository, db, validate)
	carController := controller.NewCarController(carService)

	router := app.NewRoute(carController)

	router.Run(":" + os.Getenv("APP_PORT"))
}
