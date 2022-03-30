package golang

/*
 * @lc app=leetcode.cn id=97 lang=golang
 *
 * [97] Interleaving String
 */

// @lc code=start
func isInterleave(s1 string, s2 string, s3 string) bool {
	valid := false
	if len(s3) != len(s1)+len(s2) {
		return false
	}
	var backtrack func(int, int, int, int) bool
	backtrack = func(p1, p2, p3, turn int) bool {
		if valid {
			return true
		}
		if p1 == len(s1) && p2 == len(s2) && p3 == len(s3) {
			return true
		}
		if turn == 1 {
			for i := 0; p1+i < len(s1) && p3+i < len(s3) && s1[p1+i] == s3[p3+i]; i++ {
				if !valid {
					if backtrack(p1+i+1, p2, p3+i+1, -turn) {
						valid = true
					}
				}
			}
		} else {
			for i := 0; p2+i < len(s2) && p3+i < len(s3) && s2[p2+i] == s3[p3+i]; i++ {
				if !valid {
					if backtrack(p1, p2+i+1, p3+i+1, -turn) {
						valid = true
					}
				}
			}
		}
		return valid
	}
	return backtrack(0, 0, 0, 1) || backtrack(0, 0, 0, -1)
}

// @lc code=end
