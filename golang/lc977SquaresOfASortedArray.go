package golang

/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] Squares of a Sorted Array
 */

// @lc code=start
func sortedSquares(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	i, j := 0, n-1
	for pos := n - 1; pos >= 0; pos-- {
		if vi, vj := nums[i]*nums[i], nums[j]*nums[j]; vj > vi {
			ans[pos] = vj
			j--
		} else {
			ans[pos] = vi
			i++
		}
	}
	return ans
}

// @lc code=end
