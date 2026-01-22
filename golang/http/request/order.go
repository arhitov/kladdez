package request

import "strings"

// SortDir — направление сортировки.
type SortDir string

const (
	SortAsc  SortDir = "asc"
	SortDesc SortDir = "desc"
)

// Order хранит только исходную строку формата сортировки.
//
//	Example: name,-created_at,email
type Order struct {
	Sort string `json:"sort,omitempty" key:"sort" validate:"omitempty,string"`
}

// ParseOrderMap парсит строку Sort в map[поле]направление.
// Поддерживаемый формат: "field1,-field2,+field3"
// Возвращает nil, если Sort пустая.
func (o *Order) ParseOrderMap() map[string]SortDir {
	if o.Sort == "" {
		return nil
	}

	result := make(map[string]SortDir)
	parts := strings.Split(o.Sort, ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		var (
			field string
			dir   SortDir
		)
		switch {
		case strings.HasPrefix(part, "-"):
			field = part[1:]
			dir = SortDesc
		case strings.HasPrefix(part, "+"):
			field = part[1:]
			dir = SortAsc
		default:
			field = part
			dir = SortAsc
		}

		if field != "" {
			result[field] = dir
		}
	}

	return result
}
