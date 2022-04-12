package golang

/*
 * @lc app=leetcode.cn id=806 lang=golang
 *
 * [806] Number of Lines To Write String
 */

// @lc code=start
func numberOfLines(widths []int, s string) []int {
	cur := 0
	n := len(s)
	lines, lw := 0, 0
	for cur < n {
		w := widths[s[cur]-'a']
		if lw+w <= 100 {
			lw += w
		} else {
			lines++
			lw = w
		}
		cur++
	}
	return []int{lines + 1, lw}
}

// @lc code=end
