package golang

/*
 * @lc app=leetcode.cn id=1405 lang=golang
 *
 * [1405] Longest Happy String
 */

// @lc code=start
func longestDiverseString(a int, b int, c int) string {
	ans := ""
	path := []byte{}
	var backtrack func(aCnt, bCnt, cCnt int)
	backtrack = func(aCnt, bCnt, cCnt int) {
		if len(path) > len(ans) {
			ans = string(path)
		}
		if aCnt < a {
			if len(path) < 2 || (len(path) >= 2 && string(path[len(path)-2:]) != "aa") {
				path = append(path, 'a')
				backtrack(aCnt+1, bCnt, cCnt)
				path = path[:len(path)-1]
			}
		}
		if bCnt < b {
			if len(path) < 2 || (len(path) >= 2 && string(path[len(path)-2:]) != "bb") {
				path = append(path, 'b')
				backtrack(aCnt, bCnt+1, cCnt)
				path = path[:len(path)-1]
			}
		}
		if cCnt < c {
			if len(path) < 2 || (len(path) >= 2 && string(path[len(path)-2:]) != "cc") {
				path = append(path, 'c')
				backtrack(aCnt, bCnt, cCnt+1)
				path = path[:len(path)-1]
			}
		}
	}
	backtrack(0, 0, 0)
	return ans
}

// @lc code=end
