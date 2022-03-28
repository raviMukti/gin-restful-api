package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raviMukti/gin-restful-api/helper"
	"github.com/raviMukti/gin-restful-api/model/web"
	"github.com/raviMukti/gin-restful-api/model/web/dto"
	"github.com/raviMukti/gin-restful-api/service"
)

type CarControllerImpl struct {
	CarService service.CarService
}

func NewCarController(carService service.CarService) CarController {
	return &CarControllerImpl{
		CarService: carService,
	}
}

func (controller *CarControllerImpl) Create(ctx *gin.Context) {
	var carCreateRequest *dto.CarCreateRequest = &dto.CarCreateRequest{}

	err := ctx.ShouldBindJSON(carCreateRequest)
	helper.PanicIfError(err)

	controller.CarService.Create(ctx, carCreateRequest)

	webResponse := web.WebResponse{
		Error:       "",
		ErrorDetail: "",
		Code:        http.StatusCreated,
		Message:     http.StatusText(http.StatusCreated),
		Data:        "",
	}

	ctx.JSON(http.StatusCreated, webResponse)
}

func (controller *CarControllerImpl) Update(ctx *gin.Context) {
	var carUpdateRequest *dto.CarUpdateRequest = &dto.CarUpdateRequest{}

	err := ctx.ShouldBindJSON(carUpdateRequest)
	helper.PanicIfError(err)

	carResponse := controller.CarService.Update(ctx, carUpdateRequest)

	webResponse := web.WebResponse{
		Error:       "",
		ErrorDetail: "",
		Code:        http.StatusOK,
		Message:     http.StatusText(http.StatusOK),
		Data:        carResponse,
	}

	ctx.JSON(http.StatusCreated, webResponse)
}

func (controller *CarControllerImpl) Remove(ctx *gin.Context) {
	carId := ctx.Param("carId")
	id, err := strconv.Atoi(carId)
	helper.PanicIfError(err)

	controller.CarService.Remove(ctx, id)

	webResponse := web.WebResponse{
		Error:       "",
		ErrorDetail: "",
		Code:        http.StatusOK,
		Message:     http.StatusText(http.StatusOK),
		Data:        "",
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *CarControllerImpl) GetById(ctx *gin.Context) {
	carId := ctx.Param("carId")
	id, err := strconv.Atoi(carId)
	helper.PanicIfError(err)

	carResponse := controller.CarService.GetById(ctx, id)

	webResponse := web.WebResponse{
		Error:       "",
		ErrorDetail: "",
		Code:        http.StatusOK,
		Message:     http.StatusText(http.StatusOK),
		Data:        carResponse,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *CarControllerImpl) GetAll(ctx *gin.Context) {
	carListResponse := controller.CarService.GetAll(ctx)

	webResponse := web.WebResponse{
		Error:       "",
		ErrorDetail: "",
		Code:        http.StatusOK,
		Message:     http.StatusText(http.StatusOK),
		Data:        carListResponse,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
