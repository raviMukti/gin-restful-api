package dto

type CarResponse struct {
	Id       int    `json:"id"`
	CarName  string `json:"car_name"`
	CarBrand string `json:"car_brand"`
	CarYear  string `json:"car_year"`
}
