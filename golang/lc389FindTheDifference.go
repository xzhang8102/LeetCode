package golang

/*
 * @lc app=leetcode.cn id=389 lang=golang
 *
 * [389] Find the Difference
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
	sum := 0
	for _, ch := range t {
		sum += int(ch)
	}
	for _, ch := range s {
		sum -= int(ch)
	}
	return byte(sum)
}

// @lc code=end
