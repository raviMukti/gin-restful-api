package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raviMukti/gin-restful-api/helper"
	"github.com/raviMukti/gin-restful-api/model/web"
	"github.com/raviMukti/gin-restful-api/model/web/dto"
	"github.com/raviMukti/gin-restful-api/service"
)

type BookController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Remove(ctx *gin.Context)
	GetById(ctx *gin.Context)
	GetAll(ctx *gin.Context)
}

type BookControllerImpl struct {
	BookService service.BookService
}

func NewBookController(bookService service.BookService) BookController {
	return &BookControllerImpl{
		BookService: bookService,
	}
}

// Create implements BookController
func (controller *BookControllerImpl) Create(ctx *gin.Context) {
	var bookCreateRequest *dto.BookCreateRequest = &dto.BookCreateRequest{}

	err := ctx.ShouldBindJSON(bookCreateRequest)
	helper.PanicIfError(err)

	controller.BookService.Create(ctx, bookCreateRequest)

	webResponse := web.WebResponse{
		Error:       "",
		ErrorDetail: "",
		Code:        http.StatusCreated,
		Message:     http.StatusText(http.StatusCreated),
		Data:        "",
	}

	ctx.JSON(http.StatusCreated, webResponse)
}

// GetAll implements BookController
func (controller *BookControllerImpl) GetAll(ctx *gin.Context) {
	bookListResponse := controller.BookService.GetAll(ctx)

	webResponse := web.WebResponse{
		Error:       "",
		ErrorDetail: "",
		Code:        http.StatusOK,
		Message:     http.StatusText(http.StatusOK),
		Data:        bookListResponse,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

// GetById implements BookController
func (controller *BookControllerImpl) GetById(ctx *gin.Context) {
	bookId := ctx.Param("bookId")

	bookResponse := controller.BookService.GetById(ctx, bookId)

	webResponse := web.WebResponse{
		Error:       "",
		ErrorDetail: "",
		Code:        http.StatusOK,
		Message:     http.StatusText(http.StatusOK),
		Data:        bookResponse,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

// Remove implements BookController
func (controller *BookControllerImpl) Remove(ctx *gin.Context) {
	bookId := ctx.Param("bookId")

	controller.BookService.Remove(ctx, bookId)

	webResponse := web.WebResponse{
		Error:       "",
		ErrorDetail: "",
		Code:        http.StatusOK,
		Message:     http.StatusText(http.StatusOK),
		Data:        "",
	}

	ctx.JSON(http.StatusOK, webResponse)
}

// Update implements BookController
func (controller *BookControllerImpl) Update(ctx *gin.Context) {
	var bookUpdateRequest *dto.BookUpdateRequest = &dto.BookUpdateRequest{}

	err := ctx.ShouldBindJSON(bookUpdateRequest)
	helper.PanicIfError(err)

	bookResponse := controller.BookService.Update(ctx, bookUpdateRequest)

	webResponse := web.WebResponse{
		Error:       "",
		ErrorDetail: "",
		Code:        http.StatusOK,
		Message:     http.StatusText(http.StatusOK),
		Data:        bookResponse,
	}

	ctx.JSON(http.StatusCreated, webResponse)
}
