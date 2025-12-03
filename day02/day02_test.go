package main

import "testing"

func TestIsValidIdPt1(t *testing.T) {
	tests := []struct {
		name string
		id   int
		want bool
	}{
		{"two digits, invalid", 11, false},
		{"two digits, valid", 12, true},
		{"big invalid", 1188511885, false},
		{"big valid", 1188511886, true},
		{"another big invalid", 38593859, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidIdPt1(tt.id)
			if got != tt.want {
				t.Errorf("isValidIdPt1(%d) = %v; want %v", tt.id, got, tt.want)
			}
		})
	}
}

func TestIsValidIdPt2(t *testing.T) {
	tests := []struct {
		name string
		id   int
		want bool
	}{
		{"three digits, invalid", 111, false},
		{"three digits, valid", 121, true},
		{"big invalid three times the same", 123123123, false},
		{"big invalid seven times the same", 1111111, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidIdPt2(tt.id)
			if got != tt.want {
				t.Errorf("isValidIdPt2(%d) = %v; want %v", tt.id, got, tt.want)
			}
		})
	}
}
