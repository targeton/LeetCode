package solution

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test01", args{"abcabcbb"}, 3},
		{"test02", args{"bbbbb"}, 1},
		{"test03", args{"pwwkew"}, 3},
		{"test04", args{"abcabcdb"}, 4},
		{"test05", args{"ab"}, 2},
		{"test06", args{"a"}, 1},
		{"test07", args{"aa"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
