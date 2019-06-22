package sort

func mergeSort(n []int) []int {
	res := make([]int, len(n))
	copy(res, n)
	divideAndMerge(n, res, 0, len(n)-1)
	return res
}

// Top-down
func divideAndMerge(n []int, res []int, l int, r int) {
	if l < r {
		m := (r + l) / 2
		divideAndMerge(n, res, l, m)
		divideAndMerge(n, res, m+1, r)
		merge(n, res, l, m+1, r)
	}
}

func merge(n []int, res []int, l int, m int, r int) {
	left, right := l, m
	// 拷贝上一递归已排好序的内容
	copy(n, res)
	for i := l; i <= r; i++ {
		if right > r || left < m && n[left] <= n[right] {
			res[i] = n[left]
			left++
		} else {
			res[i] = n[right]
			right++
		}
	}
}

//Bottom-up merge sublists two at one time util a single list remains
func mergeSort2(n []int) []int {
	sublists := make(map[int][]int)
	for i := 0; i < len(n); i++ {
		sublists[i] = []int{n[i]}
	}
	for len(sublists) > 1 {
		tmp := make(map[int][]int)
		j := 0
		for i := 0; i < len(sublists); i += 2 {
			left, _ := sublists[i]
			right, ok := sublists[i+1]
			if ok {
				m := merge2(left, right)
				tmp[j] = m
			} else {
				tmp[j] = left
			}
			j++
		}
		sublists = tmp
	}
	return sublists[0]
}

func merge2(left []int, right []int) []int {
	res := make([]int, len(left)+len(right))
	l, r := 0, 0
	for i := 0; i < len(res); i++ {
		if r >= len(right) || l < len(left) && left[l] <= right[r] {
			res[i] = left[l]
			l++
		} else {
			res[i] = right[r]
			r++
		}
	}
	return res
}
