package solution

/* if input is 64-bit signed integer, use the flowing code to determine integer overflow
func reverse(x int) int {
	result, newResult := 0, 0
	for {
		tail := x % 10
		newResult = result * 10 + tail
		if (newResult - tail) / 10 != result {
			return 0
		}
		result = newResult
		x = x / 10
		if x == 0{
			break
		}
	}
	return result
}
*/
func reverse(x int) int {
	result := 0
	const int32max = 1<<31 - 1
	const int32min = -1 * (1 << 31)
	for {
		result = result*10 + x%10
		x = x / 10
		if result > int32max || result < int32min {
			result = 0
			break
		}
		if x == 0 {
			break
		}
	}
	return result
}
