package arrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestIsPermutation(t *testing.T) {
	testCases := []struct {
		input1 string
		input2 string
		output bool
	}{
		{"", "", true},
		{"a", "a", true},
		{"abcA", "aAbc", true},
		{"heo", "oeh", true},
		{"hello", "fello", false},
		{"aba", "aba", true},
		{"aba", "aca", false},
		{"aaaaaaaaa", "aaa", false},
		{"blu bla", "ull abb", true},
		{"Lorem ipsum dolor sit amet.", "blah", false},
		// The above tests could be structured better by recognizing that
		// there are only 3 possible class of inputs
		// 1. Same length strings [permutation, not permutation]
		// 2. Different length strings
	}
	testFuncs := []func(string, string) bool{isPermutationSort, isPermutationMap}
	for _, fn := range testFuncs {
		for _, tc := range testCases {
			t.Run(fmt.Sprintf("Test %s & %s", tc.input1, tc.input2), func(t *testing.T) {
				assert.Equal(t, tc.output, fn(tc.input1, tc.input2))
			})
		}
	}
}
