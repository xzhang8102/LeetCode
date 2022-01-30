package golang

/*
 * @lc app=leetcode.cn id=1342 lang=golang
 *
 * [1342] Number of Steps to Reduce a Number to Zero
 */

// @lc code=start
func numberOfSteps(num int) int {
	ans := 0
	for num > 0 {
		ans++
		if num&1 == 1 {
			num--
		} else {
			num >>= 1
		}
	}
	return ans
}

// @lc code=end
