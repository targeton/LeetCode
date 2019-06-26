package solution

import (
	"reflect"
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
		{"test01", args{4}, [][]string{[]string{".Q..", "...Q", "Q...", "..Q."}, []string{"..Q.", "Q...", "...Q", ".Q.."}}},
		{"test02", args{1}, [][]string{[]string{"Q"}}},
		{"test03", args{2}, [][]string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveNQueens(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}
