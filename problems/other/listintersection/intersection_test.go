package listintersection

import (
	"reflect"
	"testing"
)

func TestListFromString(t *testing.T) {
	type args struct {
		list string
	}
	tests := []struct {
		name string
		args args
		want *List
	}{
		{
			name: "1->2->3",
			args: args{list: "1->2->3"},
			want: &List{head: &Node{
				val: "1",
				next: &Node{
					val: "2",
					next: &Node{
						val: "3",
					},
				},
			},
			},
		},
		{
			name: "1->2->",
			args: args{list: "1->2->"},
			want: &List{head: &Node{
				val: "1",
				next: &Node{
					val: "2",
					next: &Node{
						val: "",
					},
				},
			},
			},
		},
		{
			name: "",
			args: args{list: ""},
			want: &List{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListFromString(tt.args.list); got.String() != tt.want.String() {
				t.Errorf("ListFromString() = %v, want %v", got.String(), tt.want.String())
			}
		})
	}
}

func TestList_CheckIntersection(t *testing.T) {
	type args struct {
		ll *List
	}
	tests := []struct {
		name string
		list *List
		args args
		want *Node
	}{
		{
			name: "first smaller: 3->7->8->10 and 1->2->3->99->1->8->10",
			list: ListFromString("3->7->8->10"),
			args: args{ll: ListFromString("1->2->3->99->1->8->10")},
			want: ListFromString("8->10").head,
		},
		{
			name: "3->7->8->10 and 99->1->8->10",
			list: ListFromString("3->7->8->10"),
			args: args{ll: ListFromString("99->1->8->10")},
			want: ListFromString("8->10").head,
		},
		{
			name: "No intersection",
			list: ListFromString("3->7->8->10"),
			args: args{ll: ListFromString("3->4->5")},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.list.CheckIntersection(tt.args.ll); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.CheckIntersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
