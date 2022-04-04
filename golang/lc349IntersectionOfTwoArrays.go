package golang

/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	set := map[int]int{}
	for _, v := range nums1 {
		set[v] = 1
	}
	ans := []int{}
	for _, v := range nums2 {
		if set[v] > 0 {
			ans = append(ans, v)
			set[v]--
		}
	}
	return ans
}

// @lc code=end
