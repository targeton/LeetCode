package solution

func findKthLargest(nums []int, k int) int {
	return find(nums, k, 0, len(nums)-1)
}

func find(nums []int, k int, left int, right int) int {
	l, r := left, right
	for i := l + 1; i <= r; {
		if nums[i] > nums[l] {
			nums[i], nums[r] = nums[r], nums[i]
			r--
		} else {
			nums[i], nums[l] = nums[l], nums[i]
			l++
			i++
		}
	}
	if k-1 == right-r {
		return nums[r]
	} else if k-1 > right-r {
		return find(nums, k-1-(right-r), left, r-1)
	} else {
		return find(nums, k, r+1, right)
	}
}
