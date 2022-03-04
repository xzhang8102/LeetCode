package golang

/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	mapping := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	ans := 0
	for i := range s {
		if i < len(s)-1 && mapping[s[i+1]] > mapping[s[i]] {
			ans -= mapping[s[i]]
		} else {
			ans += mapping[s[i]]
		}
	}
	return ans
}

// @lc code=end
