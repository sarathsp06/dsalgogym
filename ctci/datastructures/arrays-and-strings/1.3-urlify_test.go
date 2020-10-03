package arrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLify(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty strings",
			args: args{s: ""},
			want: "",
		},
		{
			name: "simple string",
			args: args{s: "sarath"},
			want: "sarath",
		},
		{
			name: "string with spaces",
			args: args{s: "sarath sarath"},
			want: `sarath%20sarath`,
		},
		{
			name: "only spaces",
			args: args{s: "  "},
			want: `%20%20`,
		},
		{
			name: "one letter between spaces",
			args: args{s: " a "},
			want: `%20a%20`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := URLify(tt.args.s); got != tt.want {
				t.Errorf("\nURLify(%s) = [%v], want [%v]", tt.args.s, got, tt.want)
			}
		})
	}
}

func TestURLify2(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output string
	}{
		{"Empty String", "", ""},
		{"No space", "allbasdlf", "allbasdlf"},
		{"Single space", "hello world", "hello%20world"},
		{"Multiple spaces", "hello world !", "hello%20world%20!"},
		{"Multiple consecutive spacess", "hello   world", "hello%20%20%20world"},
	}
	testFuncs := []func(string) string{URLifyBuiltin}
	for _, fn := range testFuncs {
		for _, tc := range testCases {
			t.Run(fmt.Sprint(tc.name), func(t *testing.T) {
				assert.Equal(t, tc.output, fn(tc.input))
			})
		}
	}
}
