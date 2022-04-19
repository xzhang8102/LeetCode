package golang

import "strings"

/*
 * @lc app=leetcode.cn id=388 lang=golang
 *
 * [388] Longest Absolute File Path
 */

// @lc code=start
func lengthLongestPath(input string) int {
	ans := 0
	arr := strings.Split(input, "\n")
	stack := []struct{ path, lvl int }{}
	for _, s := range arr {
		name := strings.TrimLeft(s, "\t")
		level := len(s) - len(name)
		for len(stack) > 0 && stack[len(stack)-1].lvl != level-1 {
			stack = stack[:len(stack)-1]
		}
		pathL := len(name)
		if len(stack) > 0 {
			pathL += stack[len(stack)-1].path + 1
		}
		if strings.Contains(name, ".") {
			ans = lc388Max(ans, pathL)
		} else {
			stack = append(stack, struct {
				path int
				lvl  int
			}{pathL, level})
		}
	}
	return ans
}

func lc388Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
