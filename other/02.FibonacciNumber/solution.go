package solution

func fib(n int) int {
	if n < 2 {
		return n
	}
	first := 0
	second := 1
	for i := 2; i <= n; i++ {
		third := first + second
		first, second = second, third
	}
	return second
}

func climbStairs(n int) int {
	if n < 2 {
		return n
	}
	first := 1
	second := 2
	for i := 3; i <= n; i++ {
		third := first + second
		first, second = second, third
	}
	return second
}
