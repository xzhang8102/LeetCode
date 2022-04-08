package golang

/*
 * @lc app=leetcode.cn id=389 lang=golang
 *
 * [389] Find the Difference
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
	ans := byte(0)
	for _, ch := range s {
		ans ^= byte(ch)
	}
	for _, ch := range t {
		ans ^= byte(ch)
	}
	return ans
}

// @lc code=end
