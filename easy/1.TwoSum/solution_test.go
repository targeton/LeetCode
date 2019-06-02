package solution

import "testing"

func TestTwoSum(t *testing.T) {
	nums, target := []int{2, 3, 2, 7, 11, 15}, 9
	excepted := []int{0, 3}

	output := twoSum(nums, target)
	for i := range excepted {
		if output[i] != excepted[i] {
			t.Errorf("\n nums: %v, target: %v\n excepted: %v\n output: %v \n", nums, target, excepted, output)
		}
	}
}
