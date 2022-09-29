package common

var DefaultLimit = 10

type FilterOptions[T any] func(*T)

type SearchOptions[T any] struct {
	Pagination *PaginationOptions `json:"pagination"`
	Filters    []FilterOptions[T] `json:"filters"`
}

type PaginationOptions struct {
	Offset *int `json:"offset"`
	Limit  *int `json:"limit"`
}
