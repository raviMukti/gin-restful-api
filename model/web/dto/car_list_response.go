package dto

type CarListResponse struct {
	CarList      []CarResponse `json:"car_list"`
	TotalRecords int           `json:"total_records"`
}
