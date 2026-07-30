package gostring

import (
	"testing"
)

func TestPrefixBefore(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		sep      string
		expected string
	}{
		{"empty", "", "", ""},
		{"sep_empty", "qweadszx dfsnn fghtre", "", ""},
		{"none", "qweadszx dfsnn fghtre", "<GS>", "qweadszx dfsnn fghtre"},
		{"start", "<GS>qweadszx dfsnn fghtre", "<GS>", ""},
		{"center", "qweadszx <GS>dfsnn fghtre<GS>", "<GS>", "qweadszx "},
		{"end", "qweadszx dfsnn fghtre<GS>", "<GS>", "qweadszx dfsnn fghtre"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PrefixBefore(tt.input, tt.sep)
			if result != tt.expected {
				t.Errorf("PrefixBefore(%s, %s) = %s, expected %s",
					tt.input, tt.sep, result, tt.expected)
			}
		})
	}
}
