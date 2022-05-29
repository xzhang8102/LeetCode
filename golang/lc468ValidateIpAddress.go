package golang

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=468 lang=golang
 *
 * [468] Validate IP Address
 */

// @lc code=start
func validIPAddress(queryIP string) string {
	if sp := strings.Split(queryIP, "."); len(sp) == 4 {
		for _, segment := range sp {
			if len(segment) > 1 && segment[0] == '0' {
				return "Neither"
			}
			val, err := strconv.Atoi(segment)
			if err != nil || val < 0 || val > 255 {
				return "Neither"
			}
		}
		return "IPv4"
	}
	if sp := strings.Split(queryIP, ":"); len(sp) == 8 {
		for _, segment := range sp {
			if len(segment) == 0 || len(segment) > 4 {
				return "Neither"
			}
			if _, err := strconv.ParseUint(segment, 16, 64); err != nil {
				return "Neither"
			}
		}
		return "IPv6"
	}
	return "Neither"
}

// @lc code=end
