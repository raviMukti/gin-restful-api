package dto

type BookUpdateRequest struct {
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
	Year   string `json:"year" validate:"required"`
}
