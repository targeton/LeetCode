package solution

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	i, j, length := 0, 0, 0
	if (m+n)%2 == 0 {
		length = (m+n)/2 + 1
	} else {
		length = (m+n)/2 + 1
	}
	index, temp := 0, make([]int, length)
	for {
		if i >= m {
			temp[index] = nums2[j]
			j++
		} else if j >= n {
			temp[index] = nums1[i]
			i++
		} else if nums1[i] < nums2[j] {
			temp[index] = nums1[i]
			i++
		} else {
			temp[index] = nums2[j]
			j++
		}
		index++
		if index >= length {
			break
		}
	}
	result := 0.0
	if (m+n)%2 == 0 {
		result = float64(temp[length-1]+temp[length-2]) / 2.0
	} else {
		result = float64(temp[length-1])
	}
	return result
}
