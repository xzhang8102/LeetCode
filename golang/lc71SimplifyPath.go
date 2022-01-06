package golang

import "strings"

/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] Simplify Path
 */

// @lc code=start
func simplifyPath(path string) string {
	n := len(path)
	fs := []string{""}
	ptr := 0
	for i := 0; i < n; {
		start := i
		for ; start < n && path[start] != '/'; start++ {
		}
		if start > i {
			content := path[i:start]
			if content == ".." {
				if ptr > 0 {
					fs = fs[:ptr]
					ptr--
				}
			} else if content != "." {
				fs = append(fs, content)
				ptr++
			}
		}
		i = start + 1
	}
	return "/" + strings.Join(fs[1:ptr+1], "/")
}

// @lc code=end
