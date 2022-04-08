package golang

/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] First Unique Character in a String
 */

// @lc code=start
func firstUniqChar(s string) int {
	dict := [26]int{}
	for i, ch := range s {
		switch dict[ch-'a'] {
		case 0:
			dict[ch-'a'] = i + 1
		default:
			dict[ch-'a'] = -1
		}
	}
	ans := -1
	for _, v := range dict {
		if v > 0 {
			if ans == -1 || v < ans {
				ans = v
			}
		}
	}
	if ans > 0 {
		return ans - 1
	}
	return ans
}

// @lc code=end
