package arrays

import "testing"

func Test_isPermutation(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Same string",
			args: args{s1: "sarath", s2: "sarath"},
			want: true,
		},
		{
			name: "Permutation",
			args: args{s1: "test string", s2: "tste irtnsg"},
			want: true,
		},
		{
			name: "Different length",
			args: args{s1: "test s1", s2: "test s23"},
			want: false,
		},
		{
			name: "same length not permutation",
			args: args{s1: "test s1", s2: "test s2"},
			want: false,
		},
		{
			name: "empty string",
			args: args{s1: "", s2: "test s2"},
			want: false,
		},
		{
			name: "both empty strings",
			args: args{s1: "", s2: ""},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPermutation(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
