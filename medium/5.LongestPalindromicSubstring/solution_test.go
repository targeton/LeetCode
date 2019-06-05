package solution

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test01", args{"babad"}, "bab"},
		{"test02", args{"cbbd"}, "bb"},
		{"test03", args{"abcdcbd"}, "bcdcb"},
		{"test04", args{"a"}, "a"},
		{"test05", args{"abcde"}, "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
