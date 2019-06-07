package solution

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test01", args{"PAYPALISHIRING", 3}, "PAHNAPLSIIGYIR"},
		{"test02", args{"PAYPALISHIRING", 4}, "PINALSIGYAHRPI"},
		{"test03", args{"PAYPALISHIRING", 5}, "PHASIYIRPLIGAN"},
		{"test04", args{"ABCD", 2}, "ACBD"},
		{"test05", args{"ABCD", 1}, "ABCD"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
