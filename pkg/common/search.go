package common

type SearchOptions[T any] struct {
	Pagination *PaginationOptions `json:"pagination"`
	Filters    *T                 `json:"filters"`
}

type PaginationOptions struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}
