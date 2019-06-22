package sort

func sort(n []int) []int {
	//quicksort
	// quicksort1(n, 0, len(n)-1)
	// quicksort2(n, 0, len(n)-1)
	quicksort3(n)
	return n
}

//
func quicksort1(n []int, left int, right int) {
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
		quicksort1(n, left, p-1)
	}
	if right-p > 1 {
		quicksort1(n, p+1, right)
	}
}

func quicksort2(n []int, left int, right int) {
	if left >= right {
		return
	}
	divider := left
	p := right
	for i := left + 1; i <= p; {
		if n[i] >= n[divider] {
			n[i], n[p] = n[p], n[i]
			p--
		} else {
			n[i], n[divider] = n[divider], n[i]
			i++
			divider++
		}
	}
	quicksort2(n, left, divider-1)
	quicksort2(n, divider+1, right)
}

func quicksort3(n []int) {
	if len(n) <= 1 {
		return
	}
	head, tail, i := 0, len(n)-1, 1
	for head < tail {
		if n[i] > n[head] {
			n[i], n[tail] = n[tail], n[i]
			tail--
		} else {
			n[i], n[head] = n[head], n[i]
			i++
			head++
		}
	}
	quicksort3(n[:head])
	quicksort3(n[head+1:])
}
