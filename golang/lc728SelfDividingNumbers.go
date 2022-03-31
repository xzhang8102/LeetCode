package golang

/*
 * @lc app=leetcode.cn id=728 lang=golang
 *
 * [728] Self Dividing Numbers
 */

// @lc code=start
func selfDividingNumbers(left int, right int) []int {
	ans := []int{}
	for n := left; n <= right; n++ {
		num := n
		for num > 0 {
			if digit := num % 10; digit == 0 || n%digit != 0 {
				break
			}
			num /= 10
		}
		if num == 0 {
			ans = append(ans, n)
		}
	}
	return ans
}

// @lc code=end
