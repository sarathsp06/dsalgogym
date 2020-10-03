package prime

import "testing"

var dummy interface{}

func benchmarkGetMaxPrimeFact(val int64, b *testing.B) {
	var t int64
	for n := 0; n < b.N; n++ {
		t = GetMaxPrimeFact(val)
	}
	dummy = t
}

func BenchmarkGetPrimeFact600851475143(b *testing.B) { benchmarkGetMaxPrimeFact(600851475143, b) }
func BenchmarkGetPrimeFact13(b *testing.B)           { benchmarkGetMaxPrimeFact(13, b) }
func BenchmarkGetPrimeFact3326427(b *testing.B)      { benchmarkGetMaxPrimeFact(3326427, b) }

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
