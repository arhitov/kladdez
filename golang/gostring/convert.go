package gostring

import "strings"

func CaseInsensitivity(s, t string) bool {
	return strings.EqualFold(s, t)
}
