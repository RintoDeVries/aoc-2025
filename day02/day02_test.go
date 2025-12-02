package main

import "testing"

func TestIsValidId(t *testing.T) {
	tests := []struct {
		name string
		id   int
		want bool
	}{
		{"a", 11, false},
		{"b", 12, true},
		{"c", 1188511885, false},
		{"c", 1188511886, true},
		{"e", 38593859, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidId(tt.id)
			if got != tt.want {
				t.Errorf("isValidId(%d) = %v; want %v", tt.id, got, tt.want)
			}
		})
	}
}
