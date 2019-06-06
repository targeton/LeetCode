package solution

import (
	"reflect"
	"testing"
)

func Test_getNextArray(t *testing.T) {
	type args struct {
		pattern string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"test01", args{"ab"}, []int{-1, 0}},
		{"test02", args{"aa"}, []int{-1, -1}},
		{"test03", args{"abcd"}, []int{-1, 0, 0, 0}},
		{"test04", args{"abcabcd"}, []int{-1, 0, 0, -1, 0, 0, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNextArray(tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNextArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findPosition(t *testing.T) {
	type args struct {
		m string
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test01", args{"abcde", "bc"}, 1},
		{"test02", args{"abcabdc", "abd"}, 3},
		{"test03", args{"abdbbabdb", "abdba"}, -1},
		{"test04", args{"aabaaaabaaaba", "aaaba"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPosition(tt.args.m, tt.args.s); got != tt.want {
				t.Errorf("findPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}
