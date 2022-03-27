package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/raviMukti/gin-restful-api/helper"
	"github.com/raviMukti/gin-restful-api/model/domain"
)

type CarRepositoryImpl struct {
}

func NewCarRepository() CarRepository {
	return &CarRepositoryImpl{}
}

func (repository *CarRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, car domain.Car) domain.Car {
	Query := "INSERT INTO car(car_name, car_brand, car_year) VALUES (?, ?, ?)"
	result, err := tx.ExecContext(ctx, Query, car.CarName, car.CarBrand, car.CarYear)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	car.Id = id
	return car
}

func (repository *CarRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, car domain.Car) domain.Car {
	Query := "UPDATE car SET car_name = ?, car_brand = ?, car_year = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, Query, car.CarName, car.CarBrand, car.CarYear, car.Id)
	helper.PanicIfError(err)

	return car
}

func (repository *CarRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, car domain.Car) {
	Query := "DELETE FROM car WHERE id = ?"
	_, err := tx.ExecContext(ctx, Query, car.Id)
	helper.PanicIfError(err)
}

func (repository *CarRepositoryImpl) FindBydId(ctx context.Context, tx *sql.Tx, carId int) (domain.Car, error) {
	Query := "SELECT id, car_name, car_brand, car_year FROM car WHERE id = ?"
	rows, err := tx.QueryContext(ctx, Query, carId)
	helper.PanicIfError(err)
	defer rows.Close()

	car := domain.Car{}

	if rows.Next() {
		err := rows.Scan(&car.Id, &car.CarName, &car.CarBrand, &car.CarYear)
		helper.PanicIfError(err)
		return car, nil
	} else {
		return car, errors.New("car not found")
	}
}

func (repository *CarRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Car {
	Query := "SELECT id, car_name, car_brand, car_year FROM car"
	rows, err := tx.QueryContext(ctx, Query)
	helper.PanicIfError(err)
	defer rows.Close()

	var cars []domain.Car

	for rows.Next() {
		car := domain.Car{}
		err := rows.Scan(&car.Id, &car.CarName, &car.CarBrand, &car.CarYear)
		helper.PanicIfError(err)
		cars = append(cars, car)
	}

	return cars
}
