package solution

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{&ListNode{3, &ListNode{4, &ListNode{5, nil}}}, &ListNode{2, &ListNode{6, &ListNode{7, nil}}}},
			want: &ListNode{5, &ListNode{0, &ListNode{3, &ListNode{1, nil}}}},
		},
		{
			name: "test02",
			args: args{&ListNode{8, &ListNode{6, &ListNode{5, &ListNode{1, nil}}}}, &ListNode{2, &ListNode{6, &ListNode{7, nil}}}},
			want: &ListNode{0, &ListNode{3, &ListNode{3, &ListNode{2, nil}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
