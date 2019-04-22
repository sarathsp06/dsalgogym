package arrays

import "testing"

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := URLify(tt.args.s); got != tt.want {
				t.Errorf("\nURLify(%s) = [%v], want [%v]", tt.args.s, got, tt.want)
			}
		})
	}
}
