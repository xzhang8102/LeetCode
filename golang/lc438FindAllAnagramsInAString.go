package golang

/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] Find All Anagrams in a String
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	sLen, pLen := len(s), len(p)
	ans := []int{}
	window := [26]int{}
	for _, c := range p {
		window[c-'a']++
	}
	for lo, hi := 0, 0; hi < sLen; {
		if window[s[hi]-'a'] > 0 {
			window[s[hi]-'a']--
			hi++
			if hi-lo == pLen {
				ans = append(ans, lo)
			}
		} else {
			window[s[lo]-'a']++
			lo++
		}
	}
	return ans
}

// @lc code=end
