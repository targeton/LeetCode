package solution

func convert(s string, numRows int) string {
	result := ""
	length := len(s)
	pos := make([]int, length)
	if numRows > 1 {
		for i := range s {
			q := i % (2 * (numRows - 1))
			if q > numRows-1 {
				q = 2*(numRows-1) - q
			}
			pos[i] = q
		}
	}
	for i := 0; i < numRows; i++ {
		for j := range s {
			if i == pos[j] {
				result += string(s[j])
			}
		}
	}
	return result
}
