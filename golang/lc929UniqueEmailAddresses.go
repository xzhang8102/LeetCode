package golang

import "strings"

/*
 * @lc app=leetcode.cn id=929 lang=golang
 *
 * [929] Unique Email Addresses
 */

// @lc code=start
func numUniqueEmails(emails []string) int {
	set := map[string]bool{}
	for _, email := range emails {
		sp := strings.Split(email, "@")
		local := sp[0]
		plus := strings.IndexByte(local, '+')
		if plus != -1 {
			local = local[:plus]
		}
		local = strings.ReplaceAll(local, ".", "")
		set[local+"@"+sp[1]] = true
	}
	return len(set)
}

// @lc code=end
