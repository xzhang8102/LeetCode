package golang

/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start
import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	m, n := len(nums1), len(nums2)
	// [0, m]区间内进行二分查找
	lo, hi := 0, m
	halfSize := (m + n) >> 1
	odd := (m+n)%2 == 1
	for lo <= hi {
		mid := (lo + hi) >> 1
		nums1L, nums2L := math.MinInt32, math.MinInt32
		if mid > 0 {
			nums1L = nums1[mid-1]
		}
		if halfSize-mid > 0 {
			nums2L = nums2[halfSize-mid-1]
		}
		nums1R, nums2R := math.MaxInt32, math.MaxInt32
		if mid < m {
			nums1R = nums1[mid]
		}
		if halfSize-mid < n {
			nums2R = nums2[halfSize-mid]
		}
		if nums1L <= nums2R && nums2L <= nums1R {
			if odd {
				return float64(lc4Min(nums1R, nums2R))
			} else {
				return (float64(lc4Max(nums1L, nums2L)) + float64(lc4Min(nums1R, nums2R))) / 2.0
			}
		} else if nums1L >= nums2R {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return 0.0
}

func lc4Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func lc4Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
