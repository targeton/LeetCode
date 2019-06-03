package solution

func lengthOfLongestSubstring(s string) int {
	i, j, result, length := 0, 0, 0, len(s)
	for i = 0; i < length; i++ {
		cache := make(map[byte]byte)
		cache[s[i]] = 1
		for j = i + 1; j < length; j++ {
			if _, ok := cache[s[j]]; ok {
				break
			}
			cache[s[j]] = 1
		}
		if result < (j - i) {
			result = j - i
		}
	}
	return result
}
