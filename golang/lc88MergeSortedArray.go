package golang

/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	ptr, ptr1, ptr2 := m+n-1, m-1, n-1
	for ; ptr >= 0 && ptr1 >= 0 && ptr2 >= 0; ptr-- {
		if nums1[ptr1] >= nums2[ptr2] {
			nums1[ptr] = nums1[ptr1]
			ptr1--
		} else {
			nums1[ptr] = nums2[ptr2]
			ptr2--
		}
	}
	if ptr2 >= 0 {
		copy(nums1[:ptr+1], nums2[:ptr2+1])
	}
}

// @lc code=end
