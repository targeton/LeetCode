package solution

import (
	"reflect"
	"testing"
)

func Test_dailyTemperatures(t *testing.T) {
	type args struct {
		T []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"test01", args{[]int{73, 74, 75, 71, 69, 72, 76, 73}}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{"test02", args{[]int{1, 2, 3, 4, 5, 6, 7, 8}}, []int{1, 1, 1, 1, 1, 1, 1, 0}},
		{"test03", args{[]int{1}}, []int{0}},
		{"test04", args{[]int{1, 3, 5, 7, 2, 4, 6, 8}}, []int{1, 1, 1, 4, 1, 1, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dailyTemperatures(tt.args.T); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dailyTemperatures() = %v, want %v", got, tt.want)
			}
		})
	}
}
