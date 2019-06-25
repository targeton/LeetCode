package solution

func search(nums []int, target int) int {
	return searchNum(nums, target, 0, len(nums)-1)
}

/*
[5,6,7,8,9,10,11,0,1,2,3]
[8,9,10,11,0,1,2,3,4,5,6]
*/
func searchNum(nums []int, target int, head int, tail int) int {
	if head > tail {
		return -1
	}
	mid := (head + tail) / 2
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		if nums[tail] < nums[mid] && nums[head] > target {
			return searchNum(nums, target, mid+1, tail)
		}
		return searchNum(nums, target, head, mid-1)
	} else {
		if nums[head] > nums[mid] && nums[tail] < target {
			return searchNum(nums, target, head, mid-1)
		}
		return searchNum(nums, target, mid+1, tail)
	}
}
