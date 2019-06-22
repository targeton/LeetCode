package sort

import (
	"reflect"
	"testing"
)

func Test_mergeSort(t *testing.T) {
	type args struct {
		n []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"test01", args{[]int{1}}, []int{1}},
		{"test02", args{[]int{2, 3, 1}}, []int{1, 2, 3}},
		{"test03", args{[]int{5, 6, 4, 9, 3, 2, 7, 1, 8}}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{"test04", args{[]int{0, 1, 0, 4, 0, 2, 3}}, []int{0, 0, 0, 1, 2, 3, 4}},
		{"test05", args{[]int{1, 2, 1, 3, 2, 3, 4}}, []int{1, 1, 2, 2, 3, 3, 4}},
		{"test06", args{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeSort2(t *testing.T) {
	type args struct {
		n []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		// {"test01", args{[]int{1}}, []int{1}},
		// {"test02", args{[]int{2, 3, 1}}, []int{1, 2, 3}},
		{"test03", args{[]int{5, 6, 4, 9, 3, 2, 7, 1, 8}}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{"test04", args{[]int{0, 1, 0, 4, 0, 2, 3}}, []int{0, 0, 0, 1, 2, 3, 4}},
		{"test05", args{[]int{1, 2, 1, 3, 2, 3, 4}}, []int{1, 1, 2, 2, 3, 3, 4}},
		{"test06", args{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort2(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort2() = %v, want %v", got, tt.want)
			}
		})
	}
}
