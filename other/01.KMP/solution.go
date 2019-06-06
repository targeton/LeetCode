package solution

/* note:the length of pattern not less than 2
KMP算法核心就是求取匹配字符串的回退数组next []int
当匹配字符串s的s[j]与主字符串m第i位字符不匹配时（即s[0:j]与m[i-j:i]字符相同），则j取next[j]值，重新与主字符串第i位比较
不必每次都从头开始比较，充分运用匹配字符串中含有的特征
*/
func getNextArray(pattern string) []int {
	result := make([]int, len(pattern))
	result[0] = -1
	k, j := -1, 0
	for {
		if j >= len(pattern)-1 {
			break
		}
		if k == -1 || pattern[j] == pattern[k] {
			j++
			k++
			if pattern[j] == pattern[k] {
				result[j] = result[k]
			} else {
				result[j] = k
			}
		} else {
			k = result[k]
		}
	}
	return result
}

// 找出字符串m中最早匹配模式字符串s的位置，如无匹配则返回-1，利用KMP算法
func findPosition(m string, s string) int {
	i, j, pos := 0, 0, -1
	next := getNextArray(s)
	for {
		if i >= len(m) || j >= len(s) {
			break
		}
		if j == -1 || m[i] == s[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j >= len(s) {
		pos = i - len(s)
	}
	return pos
}
