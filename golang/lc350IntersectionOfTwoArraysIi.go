package golang

/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] Intersection of Two Arrays II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		// reduce the size of `set`
		return intersect(nums2, nums1)
	}
	set := map[int]int{}
	for _, v := range nums1 {
		set[v]++
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
