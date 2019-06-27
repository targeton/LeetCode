package solution

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test00", args{"aa", "a"}, false},
		{"test01", args{"aa", "a*"}, true},
		{"test02", args{"ab", ".*"}, true},
		{"test03", args{"aab", "c*a*b"}, true},
		{"test04", args{"mississippi", "mis*is*p*."}, false},
		{"test05", args{"", ""}, true},
		{"test06", args{"", "a*"}, true},
		{"test07", args{"c", "."}, true},
		{"test08", args{"cb", ".a"}, false},
		{"test09", args{"ab", ".*c"}, false},
		{"test10", args{"aaa", "a*a"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
