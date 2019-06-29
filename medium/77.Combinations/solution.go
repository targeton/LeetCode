package solution

import "math/bits"

func combine(n int, k int) [][]int {
	ret := make([][]int, 0)
	// 1 << uint(n) 表示二进制后n位为组合置位区间
	for i := 0; i < (1 << uint(n)); i++ {
		// 用bit位置位来表示组合，例如 3 二进制 0000 0101, 即有两位置1, 循环出哪两位置位，即得到这次组合是{2,4}
		if bits.OnesCount32(uint32(i)) == k {
			var cur []int
			for j := 0; j < n; j++ {
				if i&(1<<uint(j)) != 0 {
					cur = append(cur, j+1)
				}
			}
			ret = append(ret, cur)
		}
	}
	return ret
}
