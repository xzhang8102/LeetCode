package golang

/*
 * @lc app=leetcode.cn id=780 lang=golang
 *
 * [780] Reaching Points
 */

// @lc code=start
func reachingPoints(sx int, sy int, tx int, ty int) bool {
	for tx > sx && ty > sy && tx != ty {
		if tx > ty {
			tx %= ty
		} else {
			ty %= tx
		}
	}
	if tx == sx && ty == sy {
		return true
	}
	if tx == sx {
		return ty > sy && (ty-sy)%sx == 0
	}
	if ty == sy {
		return tx > sx && (tx-sx)%sy == 0
	}
	return false
}

// @lc code=end
