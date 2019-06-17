package solution

func dailyTemperatures(T []int) []int {
	length := len(T)
	result := make([]int, length)
	for i := length - 2; i >= 0; i-- {
		k := i + 1
		for k < length {
			if T[i] < T[k] {
				result[i] = k - i
				break
			} else if result[k] == 0 {
				break
			} else {
				k = k + result[k]
			}
		}
	}
	return result
}
