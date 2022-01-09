package golang

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start
func letterCombinations(digits string) []string {
	ans := []string{}
	if len(digits) == 0 {
		return ans
	}
	mapping := [10]string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	pick := []byte{}
	var dfs func(int)
	dfs = func(index int) {
		if len(pick) == len(digits) {
			ans = append(ans, string(append([]byte(nil), pick...)))
			return
		}
		for _, char := range mapping[digits[index]-'0'] {
			pick = append(pick, byte(char))
			dfs(index + 1)
			pick = pick[:len(pick)-1]
		}
	}
	dfs(0)
	return ans
}

// @lc code=end
