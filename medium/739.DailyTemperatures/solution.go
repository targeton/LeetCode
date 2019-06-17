package solution

func dailyTemperatures(T []int) []int {
	length := len(T)
	result := make([]int, length)
	stack := make([]int, length)
	top := -1
	for i := 0; i < length; i++ {
		for top > -1 && T[i] > T[stack[top]] {
			idx := stack[top]
			result[idx] = i - idx
			top--
		}
		top++
		stack[top] = i
	}
	return result
}
