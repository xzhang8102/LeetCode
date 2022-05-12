package golang

/*
 * @lc app=leetcode.cn id=944 lang=golang
 *
 * [944] Delete Columns to Make Sorted
 */

// @lc code=start
func minDeletionSize(strs []string) int {
	n, m := len(strs), len(strs[0])
	ans := 0
	for j := 0; j < m; j++ {
		for i := 1; i < n; i++ {
			if strs[i][j] < strs[i-1][j] {
				ans++
				break
			}
		}
	}
	return ans
}

// @lc code=end
