package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/raviMukti/gin-restful-api/exception"
	"github.com/raviMukti/gin-restful-api/helper"
	"github.com/raviMukti/gin-restful-api/model/domain"
	"github.com/raviMukti/gin-restful-api/model/web/dto"
	"github.com/raviMukti/gin-restful-api/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookService interface {
	Create(ctx *gin.Context, request *dto.BookCreateRequest)
	Update(ctx *gin.Context, request *dto.BookUpdateRequest) dto.BookResponse
	Remove(ctx *gin.Context, carId string)
	GetById(ctx *gin.Context, carId string) dto.BookResponse
	GetAll(ctx *gin.Context) dto.BookListResponse
}

type BookServiceImpl struct {
	BookRepository repository.BookRepository
	Validate       *validator.Validate
}

func NewBookService(bookRepository repository.BookRepository, validate *validator.Validate) BookService {
	return &BookServiceImpl{
		BookRepository: bookRepository,
		Validate:       validate,
	}
}

// Create implements BookService
func (service *BookServiceImpl) Create(ctx *gin.Context, request *dto.BookCreateRequest) {
	// Validate Request Body
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	book := domain.Book{
		Id:     primitive.NewObjectID(),
		Title:  request.Title,
		Author: request.Author,
		Year:   request.Year,
	}

	service.BookRepository.Save(ctx, book)
}

// GetAll implements BookService
func (service *BookServiceImpl) GetAll(ctx *gin.Context) dto.BookListResponse {

	var bookResponses []dto.BookResponse

	books := service.BookRepository.FindAll(ctx)

	for _, book := range books {
		stringObjectId := book.Id.Hex()
		book := dto.BookResponse{
			Id:     stringObjectId,
			Title:  book.Title,
			Author: book.Author,
			Year:   book.Year,
		}
		bookResponses = append(bookResponses, book)
	}

	bookListResponses := dto.BookListResponse{
		BookList:     bookResponses,
		TotalRecords: len(books),
	}

	return bookListResponses
}

// GetById implements BookService
func (service *BookServiceImpl) GetById(ctx *gin.Context, bookId string) dto.BookResponse {

	book, err := service.BookRepository.FindBydId(ctx, bookId)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	stringObjectId := book.Id.Hex()

	bookResponse := dto.BookResponse{
		Id:     stringObjectId,
		Title:  book.Title,
		Author: book.Author,
		Year:   book.Year,
	}

	return bookResponse
}

// Update implements BookService
func (service *BookServiceImpl) Update(ctx *gin.Context, request *dto.BookUpdateRequest) dto.BookResponse {
	// Validate Request Body
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	id := ctx.Param("bookId")
	objId, err := primitive.ObjectIDFromHex(id)
	helper.PanicIfError(err)

	book := domain.Book{
		Id:     objId,
		Title:  request.Title,
		Author: request.Author,
		Year:   request.Year,
	}

	book = service.BookRepository.Update(ctx, book)

	stringObjectId := book.Id.Hex()

	bookResponse := dto.BookResponse{
		Id:     stringObjectId,
		Title:  book.Title,
		Author: book.Author,
		Year:   book.Year,
	}

	return bookResponse
}

// Remove implements BookService
func (service *BookServiceImpl) Remove(ctx *gin.Context, bookId string) {
	err := service.BookRepository.Delete(ctx, bookId)
	helper.PanicIfError(err)
}
