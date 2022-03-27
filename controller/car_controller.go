package controller

import (
	"github.com/gin-gonic/gin"
)

type CarController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Remove(ctx *gin.Context)
	GetById(ctx *gin.Context)
	GetAll(ctx *gin.Context)
}
