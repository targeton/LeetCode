package solution

import "testing"

func Test_totalNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test00", args{1}, 1},
		{"test01", args{2}, 0},
		{"test02", args{3}, 0},
		{"test03", args{4}, 2},
		{"test04", args{5}, 10},
		{"test05", args{6}, 4},
		{"test06", args{7}, 40},
		{"test06", args{8}, 92},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalNQueens(tt.args.n); got != tt.want {
				t.Errorf("totalNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}
