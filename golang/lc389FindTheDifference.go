package golang

/*
 * @lc app=leetcode.cn id=389 lang=golang
 *
 * [389] Find the Difference
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
	dict := [26]int{}
	for _, ch := range s {
		dict[ch-'a']++
	}
	for _, ch := range t {
		if dict[ch-'a'] > 0 {
			dict[ch-'a']--
		} else {
			return byte(ch)
		}
	}
	return 0
}

// @lc code=end
