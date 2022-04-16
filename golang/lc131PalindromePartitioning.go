package golang

/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] Palindrome Partitioning
 */

// @lc code=start
func partition(s string) [][]string {
	ans := [][]string{}
	n := len(s)
	pick := []string{}
	var dfs func(start int)
	dfs = func(start int) {
		if start == n {
			tmp := make([]string, len(pick))
			copy(tmp, pick)
			ans = append(ans, tmp)
			return
		}
		for length := 1; length <= n-start; length++ {
			valid := true
			for i := start; i < start+length/2; i++ {
				if s[i] != s[start+length-1-(i-start)] {
					valid = false
					break
				}
			}
			if valid {
				pick = append(pick, s[start:start+length])
				dfs(start + length)
				pick = pick[:len(pick)-1]
			}
		}
	}
	dfs(0)
	return ans
}

// @lc code=end
