package solution

var fibCache = make(map[int]int)
var cache = make(map[int]int)

func fib(n int) int {
	if old, ok := fibCache[n]; ok {
		return old
	}
	if n < 2 {
		fibCache[n] = n
		return n
	}
	result := fib(n-1) + fib(n-2)
	fibCache[n] = result
	return result
}

func climbStairs(n int) int {
	if old, ok := cache[n]; ok {
		return old
	}
	if n <= 2 {
		cache[n] = n
		return n
	}
	result := climbStairs(n-1) + climbStairs(n-2)
	cache[n] = result
	return result
}
