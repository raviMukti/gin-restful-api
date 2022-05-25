package app

import (
	"github.com/gin-gonic/gin"
	"github.com/raviMukti/gin-restful-api/controller"
	"github.com/raviMukti/gin-restful-api/exception"
)

func NewRoute(carController controller.CarController) *gin.Engine {

	router := gin.Default()
	// Register Panic Handler
	router.Use(gin.CustomRecovery(exception.ErrorHandler))

	// Register Routes
	router.GET("/api/cars", carController.GetAll)
	router.GET("/api/cars/:carId", carController.GetById)
	router.POST("/api/cars", carController.Create)
	router.PUT("/api/cars/:carId", carController.Update)
	router.DELETE("/api/cars/:carId", carController.Remove)

	return router
}