package golang

/*
 * @lc app=leetcode.cn id=409 lang=golang
 *
 * [409] Longest Palindrome
 */

// @lc code=start
func longestPalindrome(s string) int {
	freq := [128]int{}
	for _, ch := range s {
		freq[ch]++
	}
	ans := 0
	odd := false
	for _, v := range freq {
		if v&1 == 0 {
			ans += v
		} else {
			if odd {
				ans += v - 1
			} else {
				ans += v
				odd = true
			}
		}
	}
	return ans
}

// @lc code=end
