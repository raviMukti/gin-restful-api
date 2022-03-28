package service

import (
	"github.com/gin-gonic/gin"
	"github.com/raviMukti/gin-restful-api/model/web/dto"
)

type CarService interface {
	Create(ctx *gin.Context, request *dto.CarCreateRequest)
	Update(ctx *gin.Context, request *dto.CarUpdateRequest) dto.CarResponse
	Remove(ctx *gin.Context, carId int)
	GetById(ctx *gin.Context, carId int) dto.CarResponse
	GetAll(ctx *gin.Context) dto.CarListResponse
}
