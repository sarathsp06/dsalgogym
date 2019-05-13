package prime

import "testing"

func TestGetMaxPrimeFact(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "General usecase",
			args: args{n: 600851475143},
			want: 6857,
		},
		{
			name: "Result less than square root",
			args: args{n: 3326427},
			want: 13,
		},
		{
			name: "Prime Number",
			args: args{n: 13},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMaxPrimeFact(tt.args.n); got != tt.want {
				t.Errorf("GetMaxPrimeFact() = %v, want %v", got, tt.want)
			}
		})
	}
}
