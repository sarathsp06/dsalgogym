package idleserver

import (
	"reflect"
	"testing"
)

func Test_upsert(t *testing.T) {
	type args struct {
		merged []period
		item   period
	}
	tests := []struct {
		name string
		args args
		want []period
	}{
		{
			name: "Joined:item not  a sub period",
			args: args{merged: []period{{1, 2}, {3, 10}}, item: period{5, 13}},
			want: []period{{1, 2}, {3, 13}},
		},
		{
			name: "Joined:item sub period",
			args: args{merged: []period{{1, 2}, {3, 10}}, item: period{5, 7}},
			want: []period{{1, 2}, {3, 10}},
		},
		{
			name: "Not Joined",
			args: args{merged: []period{{1, 2}, {3, 10}}, item: period{15, 17}},
			want: []period{{1, 2}, {3, 10}, {15, 17}},
		},
		{
			name: "Empty merged",
			args: args{merged: []period{}, item: period{15, 17}},
			want: []period{{15, 17}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := upsert(tt.args.merged, tt.args.item); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("upsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergePeriods(t *testing.T) {
	type args struct {
		p1 []period
		p2 []period
	}
	tests := []struct {
		name string
		args args
		want []period
	}{
		{
			name: "normal case",
			args: args{
				p1: []period{{1, 2}, {4, 7}, {11, 13}, {15, 21}},
				p2: []period{{3, 4}, {5, 12}, {23, 25}},
			},
			want: []period{{1, 2}, {3, 13}, {15, 21}, {23, 25}},
		},
		{
			name: "no item in p1 and p2",
			args: args{
				p1: []period{},
				p2: []period{},
			},
			want: nil,
		},
		{
			name: "one item each",
			args: args{
				p1: []period{{2, 10}},
				p2: []period{{3, 25}},
			},
			want: []period{{2, 25}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergePeriods(tt.args.p1, tt.args.p2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergePeriods() = %v, want %v", got, tt.want)
			}
		})
	}
}
