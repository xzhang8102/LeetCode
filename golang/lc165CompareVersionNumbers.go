package golang

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=165 lang=golang
 *
 * [165] Compare Version Numbers
 */

// @lc code=start
func compareVersion(version1 string, version2 string) int {
	v1, v2 := lc165Parse(version1), lc165Parse(version2)
	i, j := 0, 0
	for i < len(v1) && j < len(v2) {
		if v1[i] < v2[j] {
			return -1
		}
		if v1[i] > v2[j] {
			return 1
		}
		i, j = i+1, j+1
	}
	if i < len(v1) {
		for _, v := range v1[i:] {
			if v != 0 {
				return 1
			}
		}
	}
	if j < len(v2) {
		for _, v := range v2[j:] {
			if v != 0 {
				return -1
			}
		}
	}
	return 0
}

func lc165Parse(version string) []int {
	ans := []int{}
	var dot int
	for {
		dot = strings.IndexByte(version, '.')
		var val int
		if dot == -1 {
			val, _ = strconv.Atoi(version)
			ans = append(ans, val)
			break
		}
		val, _ = strconv.Atoi(version[:dot])
		ans = append(ans, val)
		version = version[dot+1:]
	}
	return ans
}

// @lc code=end
