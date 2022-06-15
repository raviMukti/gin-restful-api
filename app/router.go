package app

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/raviMukti/gin-restful-api/controller"
	"github.com/raviMukti/gin-restful-api/exception"
	"github.com/raviMukti/gin-restful-api/helper"
	"github.com/raviMukti/gin-restful-api/model/web"
	"github.com/raviMukti/gin-restful-api/repository"
	"github.com/raviMukti/gin-restful-api/service"
)

func Init() {
	err := godotenv.Load()
	helper.PanicIfError(err)
}

var (
	db            *sql.DB                  = InitDatabase()
	carRepository repository.CarRepository = repository.NewCarRepository()
	validate      *validator.Validate      = validator.New()
	carService    service.CarService       = service.NewCarService(carRepository, db, validate)
	carController controller.CarController = controller.NewCarController(carService)
)

func SetupRouter() *gin.Engine {

	router := gin.Default()
	router.HandleMethodNotAllowed = true
	// Register Panic Handler
	router.Use(gin.CustomRecovery(exception.ErrorHandler))
	// Register No Route Handler
	router.NoRoute(noRoutetoHost)
	// Register Method Not Allowed Handler
	router.NoMethod(noMethodAllowed)

	// Register Car Routes Group
	cars := router.Group("/api/cars")
	{
		cars.GET("/", carController.GetAll)
		cars.GET("/:carId", carController.GetById)
		cars.POST("/", carController.Create)
		cars.PUT("/:carId", carController.Update)
		cars.DELETE("/:carId", carController.Remove)
	}

	return router
}

func noRoutetoHost(ctx *gin.Context) {
	ctx.Writer.Header().Add("Content-Type", "application/json")
	ctx.Writer.WriteHeader(http.StatusNotFound)

	webResponse := web.WebResponse{
		Error:       "API_ERROR",
		ErrorDetail: ctx.Request.URL.Path + " NOT FOUND",
		Code:        http.StatusNotFound,
		Message:     http.StatusText(http.StatusNotFound),
		Data:        "",
	}

	ctx.JSON(http.StatusNotFound, webResponse)
}

func noMethodAllowed(ctx *gin.Context) {
	ctx.Writer.Header().Add("Content-Type", "application/json")
	ctx.Writer.WriteHeader(http.StatusMethodNotAllowed)

	webResponse := web.WebResponse{
		Error:       "API_ERROR",
		ErrorDetail: ctx.Request.Method + " NOT ALLOWED",
		Code:        http.StatusMethodNotAllowed,
		Message:     http.StatusText(http.StatusMethodNotAllowed),
		Data:        "",
	}

	ctx.JSON(http.StatusMethodNotAllowed, webResponse)
}
