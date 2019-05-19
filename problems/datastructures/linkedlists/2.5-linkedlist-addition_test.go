package linkedlists

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func list(nums ...int) *List {
	l := new(List)
	for _, n := range nums {
		digit := &Digit{Digit: n}
		digit.Next = l.Head
		l.Head = digit
	}
	return l
}
func TestListLen(t *testing.T) {
	type fields struct {
		Head *Digit
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "empty",
			fields: fields{Head: nil},
			want:   0,
		},
		{
			name:   "1",
			fields: fields{Head: &Digit{Digit: 1, Next: nil}},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := List{
				Head: tt.fields.Head,
			}
			if got := l.Len(); got != tt.want {
				t.Errorf("List.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListString(t *testing.T) {
	tests := []struct {
		name   string
		fields *List
		want   string
	}{
		{
			name:   "123",
			fields: list(1, 2, 3),
			want:   "123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := List{
				Head: tt.fields.Head,
			}
			if got := l.String(); got != tt.want {
				t.Errorf("List.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListAdd(t *testing.T) {
	type args struct {
		v *List
	}
	tests := []struct {
		name string
		list *List
		args args
		want string
	}{
		{
			name: "123+234",
			list: list(1, 2, 3),
			args: args{v: list(2, 3, 4)},
			want: "357",
		},
		{
			name: "123+239",
			list: list(1, 2, 3),
			args: args{v: list(2, 3, 9)},
			want: "362",
		},
		{
			name: "123+9",
			list: list(1, 2, 3),
			args: args{v: list(9)},
			want: "132",
		},
		{
			name: "9+123",
			list: list(9),
			args: args{v: list(1, 2, 3)},
			want: "132",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Print(tt.list, tt.args.v, "\n")
			tt.list.Add(tt.args.v)
			assert.Equal(t, tt.want, tt.list.String())
		})
	}
}
