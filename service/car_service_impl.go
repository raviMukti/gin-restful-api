package service

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/raviMukti/gin-restful-api/exception"
	"github.com/raviMukti/gin-restful-api/helper"
	"github.com/raviMukti/gin-restful-api/model/domain"
	"github.com/raviMukti/gin-restful-api/model/web/dto"
	"github.com/raviMukti/gin-restful-api/repository"
)

type CarServiceImpl struct {
	CarRepository repository.CarRepository
	DB            *sql.DB
	Validate      *validator.Validate
}

func NewCarService(carRepository repository.CarRepository, DB *sql.DB, validate *validator.Validate) CarService {
	return &CarServiceImpl{
		CarRepository: carRepository,
		DB:            DB,
		Validate:      validate,
	}
}

func (carService *CarServiceImpl) Create(ctx *gin.Context, request *dto.CarCreateRequest) {
	// Validate Request Body
	err := carService.Validate.Struct(request)
	helper.PanicIfError(err)

	// Begin DB Transaction
	tx, err := carService.DB.Begin()
	helper.PanicIfError(err)

	// Defer Commit or Rollback
	defer helper.CommitOrRollback(tx)

	car := domain.Car{
		CarName:  request.CarName,
		CarBrand: request.CarBrand,
		CarYear:  request.CarYear,
	}

	// Insert A Car
	car = carService.CarRepository.Save(ctx, tx, car)
}

func (carService *CarServiceImpl) Update(ctx *gin.Context, request *dto.CarUpdateRequest) dto.CarResponse {
	// Validate Request Body
	err := carService.Validate.Struct(request)
	helper.PanicIfError(err)

	// Get DB Transaction
	tx, err := carService.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	id := ctx.Param("carId")
	carId, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	car, err := carService.CarRepository.FindBydId(ctx, tx, carId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	car.CarName = request.CarName
	car.CarBrand = request.CarBrand
	car.CarYear = request.CarYear

	car = carService.CarRepository.Update(ctx, tx, car)

	carResponse := dto.CarResponse{
		Id:       int(car.Id),
		CarName:  car.CarName,
		CarBrand: car.CarBrand,
		CarYear:  car.CarYear,
	}

	return carResponse
}

func (carService *CarServiceImpl) Remove(ctx *gin.Context, carId int) {
	tx, err := carService.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	car, err := carService.CarRepository.FindBydId(ctx, tx, carId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	carService.CarRepository.Delete(ctx, tx, car)
}

func (carService *CarServiceImpl) GetById(ctx *gin.Context, carId int) dto.CarResponse {
	tx, err := carService.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	car, err := carService.CarRepository.FindBydId(ctx, tx, carId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	carResponse := dto.CarResponse{
		Id:       int(car.Id),
		CarName:  car.CarName,
		CarBrand: car.CarBrand,
		CarYear:  car.CarYear,
	}

	return carResponse
}

func (carService *CarServiceImpl) GetAll(ctx *gin.Context) dto.CarListResponse {
	tx, err := carService.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	var carResponses []dto.CarResponse

	cars := carService.CarRepository.FindAll(ctx, tx)

	for _, car := range cars {
		car := dto.CarResponse{
			Id:       int(car.Id),
			CarName:  car.CarName,
			CarBrand: car.CarBrand,
			CarYear:  car.CarYear,
		}
		carResponses = append(carResponses, car)
	}

	carListResponse := dto.CarListResponse{
		CarList:      carResponses,
		TotalRecords: len(cars),
	}

	return carListResponse
}
