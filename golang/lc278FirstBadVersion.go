package golang

/*
 * @lc app=leetcode.cn id=278 lang=golang
 *
 * [278] First Bad Version
 */

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 */
func isBadVersion(version int) bool {
	return false
}

// @lc code=start
func firstBadVersion(n int) int {
	low, hi := 1, n
	for low <= hi {
		check := (low + hi) >> 1
		if isBadVersion(check) {
			if !isBadVersion(check - 1) {
				return check
			}
			hi = check - 1
		} else {
			low = check + 1
		}
	}
	return 0
}

// @lc code=end
