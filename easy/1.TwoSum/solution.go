package solution

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)
	for i, n := range nums {
		if pos, ok := cache[target-n]; ok {
			return []int{pos, i}
		}
		if _, ok := cache[n]; !ok {
			cache[n] = i
		}
	}
	return []int{-1, -1}
}
