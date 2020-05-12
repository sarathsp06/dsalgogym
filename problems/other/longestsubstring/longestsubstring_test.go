package longestsubstring

import (
	"testing"
)

func Test_longestsubstring(t *testing.T) {
	tests := []struct {
		args string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{" ", 1},
		{"abcabcbb", 3},
	}
	for _, tt := range tests {
		t.Run("testing longestsubstring", func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args); tt.want != got {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
