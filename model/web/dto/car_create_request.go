package dto

type CarCreateRequest struct {
	CarName  string `validate:"required" json:"car_name"`
	CarBrand string `validate:"required" json:"car_brand"`
	CarYear  string `validate:"required" json:"car_year"`
}
