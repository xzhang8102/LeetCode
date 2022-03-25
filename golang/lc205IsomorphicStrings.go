package golang

/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] Isomorphic Strings
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	s2t := [256]byte{}
	t2s := [256]byte{}
	for i := range s {
		x, y := s[i], t[i]
		if s2t[x] > 0 && s2t[x] != y || t2s[y] > 0 && t2s[y] != x {
			return false
		}
		s2t[x] = y
		t2s[y] = x
	}
	return true
}

// @lc code=end
