package golang

/*
 * @lc app=leetcode.cn id=340 lang=golang
 *
 * [340] Longest Substring with At Most K Distinct Characters
 */

// @lc code=start
func lengthOfLongestSubstringKDistinct(s string, k int) int {
	n := len(s)
	if k <= 0 {
		return 0
	}
	if n <= k {
		return n
	}
	ans := 0
	window := map[byte]int{}
	for left, right := 0, 0; right < n; {
		if len(window) < k || (len(window) == k && window[s[right]] > 0) {
			window[s[right]]++
			ans = lc340Max(ans, right-left+1)
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

func lc340Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
