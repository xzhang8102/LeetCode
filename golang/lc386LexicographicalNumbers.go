package golang

/*
 * @lc app=leetcode.cn id=386 lang=golang
 *
 * [386] Lexicographical Numbers
 */

// @lc code=start
func lexicalOrder(n int) []int {
	ans := make([]int, 0, n)
	var dfs func(int)
	dfs = func(start int) {
		if start > n {
			return
		}
		bound := start + 9
		if start == 1 {
			bound = 9
		}
		for i := start; i <= bound; i++ {
			if i <= n {
				ans = append(ans, i)
				dfs(i * 10)
			} else {
				return
			}
		}
	}
	dfs(1)
	return ans
}

// @lc code=end
