package solution

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"test01", args{[]int{2, 3, 12, 5, 10, 7, 8}, 9}, []int{0, 5}},
		{"test02", args{[]int{2, 3, 12, 5, 10, 7, 8}, 22}, []int{2, 4}},
		{"test03", args{[]int{2, 3, 4, 5, 6, 7, 8, 11}, 11}, []int{3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
