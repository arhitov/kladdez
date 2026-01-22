package request

type LimitOffset struct {
	Limit  int `json:"limit,omitempty"  key:"limit"  validate:"omitempty,numeric,min=0,max=1000"`
	Offset int `json:"offset,omitempty" key:"offset" validate:"omitempty,numeric,min=0"`
}
