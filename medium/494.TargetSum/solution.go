package solution

/*
* nums can divide into P(positive subset) and N(negative subset)
* sum(P) - sum(N) = S
* sum(P) + sum(N) + sum(P) - sum(N) = S + sum(P) + sum(N)
* 2 * sum(P) = S + sum(nums)
* sum(P) = (S + sum(nums)) / 2
* the problem can be converted to a subset sum problem
 */
func findTargetSumWays(nums []int, S int) int {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	if sum < S {
		return 0
	}
	if (sum+S)%2 > 0 {
		return 0
	}
	return subSetSum(nums, S)
}

func subSetSum(nums []int, S int) int {
	dp := make([]int, S+1)
	dp[0] = 1
	for _, n := range nums {
		for j := S; j >= n; j-- {
			dp[j] += dp[j-n]
		}
	}
	return dp[S]
}
