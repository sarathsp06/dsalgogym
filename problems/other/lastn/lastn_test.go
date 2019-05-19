package lastn

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLastNRecord(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name   string
		lastN  LastN
		args   args
		expect LastN
	}{
		{
			name:   "Record first element",
			lastN:  LastN{items: []int{0, 0, 0}, last: -1, cap: 3},
			args:   args{val: 13},
			expect: LastN{items: []int{13, 0, 0}, last: 0, cap: 3},
		},
		{
			name:   "Record an element after max capacity",
			lastN:  LastN{items: []int{1, 2, 3}, last: 2, cap: 3},
			args:   args{val: 13},
			expect: LastN{items: []int{13, 2, 3}, last: 0, cap: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.lastN.Record(tt.args.val)
			assert.Equal(t, tt.expect, tt.lastN)
		})
	}
}

func TestLastNGetLast(t *testing.T) {
	type args struct {
		idx int
	}
	tests := []struct {
		name  string
		lastN LastN
		args  args
		want  int
	}{
		{
			name:  "full: last ith element",
			lastN: LastN{items: []int{1, 2, 3, 4}, cap: 4, last: 3},
			args:  args{2},
			want:  2,
		},
		{
			name:  "partial: last ith element larger than last ",
			lastN: LastN{items: []int{1, 2, 0, 0}, cap: 4, last: 1},
			args:  args{3},
			want:  0,
		},
		{
			name:  "full: recorded more entries than cap",
			lastN: LastN{items: []int{1, 2, 5, 7}, cap: 4, last: 1},
			args:  args{3},
			want:  5,
		},
		{
			name:  "empty",
			lastN: LastN{items: make([]int, 3), cap: 3, last: -1},
			args:  args{2},
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.lastN.GetLast(tt.args.idx); got != tt.want {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		cap uint
	}
	tests := []struct {
		name string
		args args
		want *LastN
	}{
		{
			name: "normal instance",
			args: args{cap: 3},
			want: &LastN{cap: 3, items: make([]int, 3)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.cap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
