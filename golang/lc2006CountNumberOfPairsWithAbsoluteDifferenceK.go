package golang

/*
 * @lc app=leetcode.cn id=2006 lang=golang
 *
 * [2006] Count Number of Pairs With Absolute Difference K
 */

// @lc code=start
func countKDifference(nums []int, k int) int {
	hashmap := map[int]int{}
	ans := 0
	for _, n := range nums {
		ans += hashmap[n-k] + hashmap[n+k]
		hashmap[n]++
	}
	return ans
}

// @lc code=end
