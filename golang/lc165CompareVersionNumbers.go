package golang

/*
 * @lc app=leetcode.cn id=165 lang=golang
 *
 * [165] Compare Version Numbers
 */

// @lc code=start
func compareVersion(version1 string, version2 string) int {
	n1, n2 := len(version1), len(version2)
	for i, j := 0, 0; i < n1 || j < n2; i, j = i+1, j+1 {
		x := 0
		for i < n1 && version1[i] != '.' {
			x = x*10 + int(version1[i]-'0')
			i++
		}
		y := 0
		for j < n2 && version2[j] != '.' {
			y = y*10 + int(version2[j]-'0')
			j++
		}
		if x < y {
			return -1
		}
		if x > y {
			return 1
		}
	}
	return 0
}

// @lc code=end
