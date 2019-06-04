package solution

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"test01", args{[]int{1, 3}, []int{2}}, 2.0},
		{"test02", args{[]int{1, 2}, []int{3, 4}}, 2.5},
		{"test03", args{[]int{1, 2, 3, 4, 5, 6, 7}, []int{3, 5, 7, 9}}, 5.0},
		{"test04", args{[]int{1, 2, 3, 4, 5, 6, 7}, []int{2, 3, 4, 7, 8, 9, 11}}, 4.5},
		{"test05", args{[]int{1}, []int{2}}, 1.5},
		{"test06", args{[]int{1}, []int{}}, 1.0},
		{"test07", args{[]int{1}, []int{2, 3}}, 2.0},
		{"test08", args{[]int{3}, []int{1, 2}}, 2.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
