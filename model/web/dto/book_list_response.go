package dto

type BookListResponse struct {
	BookList     []BookResponse `json:"book_list"`
	TotalRecords int            `json:"total_records"`
}
