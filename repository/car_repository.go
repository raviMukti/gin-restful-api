package repository

import (
	"context"
	"database/sql"

	"github.com/raviMukti/gin-restful-api/model/domain"
)

type CarRepository interface {
	Save(ctx context.Context, tx *sql.Tx, car domain.Car) domain.Car
	Update(ctx context.Context, tx *sql.Tx, car domain.Car) domain.Car
	Delete(ctx context.Context, tx *sql.Tx, car domain.Car)
	FindBydId(ctx context.Context, tx *sql.Tx, carId int) (domain.Car, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Car
}
