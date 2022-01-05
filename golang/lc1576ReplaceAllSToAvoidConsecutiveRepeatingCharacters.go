package golang

/*
 * @lc app=leetcode.cn id=1576 lang=golang
 *
 * [1576] Replace All ?'s to Avoid Consecutive Repeating Characters
 */

// @lc code=start
func modifyString(s string) string {
	ans := []byte(s)
	n := len(s)
	for i := 0; i < n; i++ {
		if ans[i] == '?' {
			for char := byte('a'); char <= 'c'; char++ {
				if !((i > 0 && char == ans[i-1]) || (i < n-1 && char == ans[i+1])) {
					ans[i] = char
					break
				}
			}
		}
	}
	return string(ans)
}

// @lc code=end
