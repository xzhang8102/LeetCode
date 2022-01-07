package golang

/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 */

// @lc code=start
func getRow(rowIndex int) []int {
	ans := make([]int, rowIndex+1)
	ans[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			ans[j] += ans[j-1]
		}
	}
	return ans
}

// @lc code=end
