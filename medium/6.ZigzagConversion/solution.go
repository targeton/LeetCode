package solution

func convert(s string, numRows int) string {
	result := ""
	if numRows == 1 {
		return s
	}
	for i := 0; i < numRows; i++ {
		for j, v := range s {
			k := j % (2 * (numRows - 1))
			if k > numRows-1 {
				k = 2*(numRows-1) - k
			}
			if i == k {
				result += string(v)
			}
		}
	}
	return result
}
