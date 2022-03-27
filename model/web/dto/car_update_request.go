package dto

type CarUpdateRequest struct {
	Id       int    `validate:"required" json:"id"`
	CarName  string `validate:"required, max=200, min=1" json:"car_name"`
	CarBrand string `validate:"required, max=200, min=1" json:"car_brand"`
	CarYear  int    `validate:"required" json:"car_year"`
}
