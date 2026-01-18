package response

import "time"

type ItemsResponse[T any] struct {
	BaseResponse
	Error      *Error       `json:"error,omitempty"`
	Data       DataItems[T] `json:"data,omitempty"`
	Pagination *Pagination  `json:"pagination,omitempty"`
	Meta       *MetaList    `json:"meta,omitempty"`
}

func NewItemsResponse[T any](
	list []T,
	total, limit, page int,
) *ItemsResponse[T] {
	return &ItemsResponse[T]{
		BaseResponse: BaseResponse{
			Status:    StatusSuccess,
			Timestamp: time.Now(),
		},
		Data: DataItems[T]{
			Items: list,
			Count: len(list),
		},
		Pagination: NewPagination(total, limit, page),
		Meta: &MetaList{
			Total: total,
			Page:  page,
			Limit: limit,
		},
	}
}

type DataItems[T any] struct {
	Items []T `json:"items"`
	Count int `json:"count"`
}

type DataItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MetaList struct {
	Total   int               `json:"total"`
	Page    int               `json:"page"`
	Limit   int               `json:"limit"`
	Sort    map[string]string `json:"sort,omitempty"`
	Filters map[string]string `json:"filters,omitempty"`
}
