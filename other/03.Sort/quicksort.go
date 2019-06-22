package sort

func sort(n []int) []int {
	//quicksort
	quicksort(n, 0, len(n)-1)
	return n
}

func quicksort(n []int, left int, right int) {
	p := left
	i, j := left, right
	for i <= j {
		for j >= p && n[j] >= n[p] {
			j--
		}
		if j >= p {
			n[p], n[j] = n[j], n[p]
			p = j
		}
		for i <= p && n[i] <= n[p] {
			i++
		}
		if i <= p {
			n[p], n[i] = n[i], n[p]
			p = i
		}
	}
	if p-left > 1 {
		quicksort(n, left, p-1)
	}
	if right-p > 1 {
		quicksort(n, p+1, right)
	}
}
