package solution

import "testing"

func Test_findTargetSumWays(t *testing.T) {
	type args struct {
		nums []int
		S    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test01", args{[]int{1, 1, 1, 1, 1}, 3}, 5},
		{"test02", args{[]int{1, 1, 1, 2, 4}, 3}, 4},
		{"test03", args{[]int{1, 1}, 3}, 0},
		{"test04", args{[]int{1, 1, 2, 4, 5}, 4}, 0},
		{"test04", args{[]int{1, 1, 2, 4, 6}, 4}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetSumWays(tt.args.nums, tt.args.S); got != tt.want {
				t.Errorf("findTargetSumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
