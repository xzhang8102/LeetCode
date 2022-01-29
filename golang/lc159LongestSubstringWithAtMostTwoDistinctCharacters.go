package golang

/*
 * @lc app=leetcode.cn id=159 lang=golang
 *
 * [159] Longest Substring with At Most Two Distinct Characters
 */

// @lc code=start
func lengthOfLongestSubstringTwoDistinct(s string) int {
	n := len(s)
	if n <= 2 {
		return n
	}
	window := map[byte]int{}
	ans := 0
	for left, right := 0, 0; right < n; {
		if len(window) < 2 || (len(window) == 2 && window[s[right]] > 0) {
			window[s[right]]++
			ans = lc159Max(ans, right-left+1)
			right++
		} else {
			if count := window[s[left]]; count == 1 {
				delete(window, s[left])
			} else {
				window[s[left]]--
			}
			left++
		}
	}
	return ans
}

func lc159Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
