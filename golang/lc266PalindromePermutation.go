package golang

/*
 * @lc app=leetcode.cn id=266 lang=golang
 *
 * [266] Palindrome Permutation
 */

// @lc code=start
func canPermutePalindrome(s string) bool {
	dict := [26]int{}
	for _, ch := range s {
		dict[ch-'a']++
	}
	cnt := 0
	for _, v := range dict {
		if v&1 == 1 {
			cnt++
		}
		if cnt > 1 {
			return false
		}
	}
	return true
}

// @lc code=end
