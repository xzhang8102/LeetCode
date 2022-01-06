package golang

import "strings"

/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] Simplify Path
 */

// @lc code=start
func simplifyPath(path string) string {
	stack := []string{}
	for _, name := range strings.Split(path, "/") {
		if name != "" {
			if name == ".." {
				if len(stack) > 0 {
					stack = stack[:len(stack)-1]
				}
			} else if name != "." {
				stack = append(stack, name)
			}
		}
	}
	return "/" + strings.Join(stack, "/")
}

// @lc code=end
