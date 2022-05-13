package golang

/*
 * @lc app=leetcode.cn id=187 lang=golang
 *
 * [187] Repeated DNA Sequences
 */

// @lc code=start
func findRepeatedDnaSequences(s string) []string {
	n := len(s)
	ans := []string{}
	set := map[string]int{}
	for i := 0; i < n-9; i++ {
		sub := s[i : i+10]
		if set[sub] == 1 {
			ans = append(ans, sub)
		}
		set[sub]++
	}
	return ans
}

// @lc code=end
