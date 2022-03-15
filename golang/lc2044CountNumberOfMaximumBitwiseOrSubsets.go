package golang

/*
 * @lc app=leetcode.cn id=2044 lang=golang
 *
 * [2044] Count Number of Maximum Bitwise-OR Subsets
 */

// @lc code=start
func countMaxOrSubsets(nums []int) int {
	maxOr := 0
	ans := 0
	var dfs func(int, int)
	dfs = func(pos, or int) {
		if pos == len(nums) {
			if or > maxOr {
				maxOr = or
				ans = 1
			} else if or == maxOr {
				ans++
			}
			return
		}
		dfs(pos+1, or|nums[pos])
		dfs(pos+1, or)
	}
	dfs(0, 0)
	return ans
}

// @lc code=end
