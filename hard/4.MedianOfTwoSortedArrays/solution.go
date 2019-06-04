package solution

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		temp := nums1
		nums1 = nums2
		nums2 = temp
	}
	m, n := len(nums1), len(nums2)
	if m == 0 {
		if n == 0 {
			return 0.0
		} else if n%2 == 0 {
			return (float64(nums2[n/2]) + float64(nums2[n/2-1])) / 2.0
		}
		return float64(nums2[n/2])
	}
	imin, imax := 0, m
	var i, j int
	maxL, minR := 0.0, 0.0
	for {
		i = (imin + imax) / 2
		j = (m+n+1)/2 - i
		if i == 0 {
			if nums1[i] >= nums2[j-1] {
				maxL = float64(nums2[j-1])
				if j < n {
					minR = math.Min(float64(nums1[i]), float64(nums2[j]))
				} else {
					minR = float64(nums1[i])
				}
				break
			} else {
				imin = i + 1
				continue
			}
		}
		if i == m {
			if nums2[j] >= nums1[i-1] {
				if j-1 >= 0 {
					maxL = math.Max(float64(nums1[i-1]), float64(nums2[j-1]))
				} else {
					maxL = float64(nums1[i-1])
				}
				minR = float64(nums2[j])
				break
			} else {
				imax = i - 1
				continue
			}
		}
		if nums1[i-1] <= nums2[j] && nums2[j-1] <= nums1[i] {
			maxL = math.Max(float64(nums1[i-1]), float64(nums2[j-1]))
			minR = math.Min(float64(nums1[i]), float64(nums2[j]))
			break
		}

		if nums1[i-1] >= nums2[j] {
			imax = i - 1
		}
		if nums2[j-1] >= nums1[i] {
			imin = i + 1
		}
	}
	if (m+n)%2 == 1 {
		return maxL
	}
	return (maxL + minR) / 2.0
}
