package request

type Pagination struct {
	PerPage int `json:"per_page,omitempty" key:"per_page" validate:"omitempty,numeric,min=0,max=1000"`
	Page    int `json:"page,omitempty"     key:"page"     validate:"omitempty,numeric,min=1"`
}
