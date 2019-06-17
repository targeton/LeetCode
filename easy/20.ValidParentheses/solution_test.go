package solution

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test01", args{"()"}, true},
		{"test02", args{"()[]{}"}, true},
		{"test03", args{"(]"}, false},
		{"test04", args{"([)]"}, false},
		{"test05", args{"{[]}"}, true},
		{"test06", args{"({[]}"}, false},
		{"test07", args{""}, true},
		{"test08", args{"]"}, false},
		{"test09", args{"()]"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
