package substr

import "testing"

func TestGetLargestSubstr(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "abcba",
			args: args{s: "abcba", k: 2},
			want: "bcb",
		},
		{
			name: "1",
			args: args{s: "1", k: 1},
			want: "1",
		},
		{
			name: "empty string",
			args: args{s: "", k: 3},
			want: "",
		},
		{
			name: "malayalam",
			args: args{s: "malayalam", k: 3},
			want: "alayala",
		},
		{
			name: "malayalam",
			args: args{s: "malayalam", k: 100},
			want: "malayalam",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLargestSubstr(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("GetLargestSubstr() = %v, want %v", got, tt.want)
			}
		})
	}
}
