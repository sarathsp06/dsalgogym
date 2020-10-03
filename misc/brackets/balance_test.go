package brackets

import "testing"

func Test_isBalanced(t *testing.T) {
	tests := []struct {
		str  string
		want bool
	}{
		{"", true},
		{"this is sparta", true},
		{"]", false},
		{"[({})]", true},
		{"[({)]", false},
		{"[()({[][]()[]})]", true},
	}
	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			if got := isBalanced(tt.str); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
