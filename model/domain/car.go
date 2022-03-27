package domain

type Car struct {
	Id       int64  `json:"id"`
	CarName  string `json:"car_name"`
	CarBrand string `json:"car_brand"`
	CarYear  int64  `json:"car_year"`
}
