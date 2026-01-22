package request

type Limit struct {
	Limit int `json:"limit,omitempty" key:"limit" validate:"omitempty,numeric,min=0,max=1000"`
}
