package golang

/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] Happy Number
 */

// @lc code=start
func isHappy(n int) bool {
	visited := map[int]bool{}
	for n != 1 && !visited[n] {
		visited[n] = true
		tmp := 0
		for n > 0 {
			tmp += (n % 10) * (n % 10)
			n /= 10
		}
		n = tmp
	}
	return n == 1
}

// @lc code=end
