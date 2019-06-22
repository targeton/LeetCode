package solution

import "testing"

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test01", args{&TreeNode{Val: 10, Left: &TreeNode{Val: 5, Left: nil, Right: nil},
			Right: &TreeNode{Val: 15, Left: &TreeNode{Val: 9, Left: nil, Right: nil},
				Right: &TreeNode{Val: 20, Left: nil, Right: nil}}}}, false},
		{"test02", args{&TreeNode{Val: 10, Left: &TreeNode{Val: 5, Left: nil, Right: nil},
			Right: &TreeNode{Val: 15, Left: &TreeNode{Val: 11, Left: nil, Right: nil},
				Right: &TreeNode{Val: 20, Left: nil, Right: nil}}}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
