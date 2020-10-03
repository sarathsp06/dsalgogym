package linkedlists

import "testing"

func TestIterativePalindrom(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{"malayalam", "MalayalaM", true},
		{"empty", "", true},
		{"paap", "paap", true},
		{"p", "p", true},
		{"thing", "thing", false},
		{"Malayalam", "Malayalam", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IterativePalindrom(tt.str); got != tt.want {
				t.Errorf("IterativePalindrom(%s) = %v, want %v", tt.str, got, tt.want)
			}
		})
	}
}
