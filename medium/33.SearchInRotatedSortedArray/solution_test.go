package solution

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test00", args{[]int{4, 5, 6, 7, 0, 1, 2}, 0}, 4},
		{"test01", args{[]int{4, 5, 6, 7, 0, 1, 2}, 3}, -1},
		{"test02", args{[]int{5, 6, 7, 8, 9, 10, 11, 0, 1, 2, 3}, 3}, 10},
		{"test03", args{[]int{5, 6, 7, 8, 9, 10, 11, 0, 1, 2, 3}, 11}, 6},
		{"test04", args{[]int{5, 6, 7, 8, 9, 10, 11, 0, 1, 2, 3}, 7}, 2},
		{"test05", args{[]int{8, 9, 10, 11, 0, 1, 2, 3, 4, 5, 6}, 3}, 7},
		{"test06", args{[]int{8, 9, 10, 11, 0, 1, 2, 3, 4, 5, 6}, 12}, -1},
		{"test07", args{[]int{8, 9, 10, 11, 0, 1, 2, 3, 4, 5, 6}, 0}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
