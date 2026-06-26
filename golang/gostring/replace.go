package gostring

import "strings"

func ReplaceAll(src string, replace []struct {
	Old string
	New string
}) string {
	var newString string = src
	for _, v := range replace {
		newString = strings.ReplaceAll(newString, v.Old, v.New)
	}
	return newString
}
