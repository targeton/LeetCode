package solution

func longestPalindrome(s string) string {
	maxLen := 0
	result := ""
	for i := range s {
		// when the length of padindrome substring is odd,s[i] is the center
		getLongestPalindrome(s, i, i, &maxLen, &result)
		// when the length of padindrome substring is even,s[i] and s[i+1] is the center
		getLongestPalindrome(s, i, i+1, &maxLen, &result)
	}
	return result
}

// get longest palindrome from center to both sides
func getLongestPalindrome(s string, left int, right int, maxLen *int, result *string) {
	for {
		if left < 0 || right >= len(s) {
			break
		}
		if s[left] == s[right] {
			if right-left+1 > *maxLen {
				*maxLen = right - left + 1
				*result = s[left : right+1]
			}
			left--
			right++
		} else {
			break
		}
	}
}
