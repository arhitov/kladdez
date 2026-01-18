package response

type Pagination struct {
	Total       int  `json:"total"`
	Page        int  `json:"page"`
	PerPage     int  `json:"per_page"`
	TotalPages  int  `json:"total_pages"`
	HasNext     bool `json:"has_next"`
	HasPrevious bool `json:"has_previous"`
}

func NewPagination(total, limit, page int) Pagination {
	totalPages := (total + limit - 1) / limit
	return Pagination{
		Total:       total,
		Page:        page,
		PerPage:     limit,
		TotalPages:  totalPages,
		HasNext:     page < totalPages,
		HasPrevious: page > 1,
	}
}
