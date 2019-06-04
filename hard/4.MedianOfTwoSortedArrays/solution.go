package solution

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m, n := len(nums1), len(nums2)
	if m == 0 {
		if n == 0 {
			return 0.0
		} else if n%2 == 0 {
			return (float64(nums2[n/2]) + float64(nums2[n/2-1])) / 2.0
		} else {
			return float64(nums2[n/2])
		}
	}
	imin, imax, half := 0, m, (m+n+1)/2
	var i, j int
	maxL, minR := 0.0, 0.0
	for {
		i = (imin + imax) / 2
		j = half - i
		if i < m && nums2[j-1] > nums1[i] {
			imin = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			imax = i - 1
		} else {
			if i == 0 {
				maxL = float64(nums2[j-1])
			} else if j == 0 {
				maxL = float64(nums1[i-1])
			} else {
				maxL = math.Max(float64(nums2[j-1]), float64(nums1[i-1]))
			}

			if i == m {
				minR = float64(nums2[j])
			} else if j == n {
				minR = float64(nums1[i])
			} else {
				minR = math.Min(float64(nums2[j]), float64(nums1[i]))
			}
			break
		}
	}
	if (m+n)%2 == 1 {
		return maxL
	}
	return (maxL + minR) / 2.0
}
