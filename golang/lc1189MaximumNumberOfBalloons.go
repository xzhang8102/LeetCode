package golang

/*
 * @lc app=leetcode.cn id=1189 lang=golang
 *
 * [1189] Maximum Number of Balloons
 */

// @lc code=start
func maxNumberOfBalloons(text string) int {
	cache := map[byte]int{}
	for i := 0; i < len(text); i++ {
		cache[text[i]]++
	}
	for i := cache['b']; i > 0; i-- {
		if cache['a'] >= i && cache['n'] >= i && cache['l'] >= 2*i && cache['o'] >= 2*i {
			return i
		}
	}
	return 0
}

// @lc code=end
