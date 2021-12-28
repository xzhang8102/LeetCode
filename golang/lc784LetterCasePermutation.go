package golang

import "strings"

/*
 * @lc app=leetcode.cn id=784 lang=golang
 *
 * [784] Letter Case Permutation
 */

// @lc code=start
func letterCasePermutation(s string) []string {
	var result []string
	var builder strings.Builder
	builder.Grow(len(s))
	lc784Helper(0, s, &builder, &result)
	return result
}

func lc784Helper(index int, s string, builder *strings.Builder, result *[]string) {
	if index == len(s) {
		*result = append(*result, builder.String())
		return
	}
	char := s[index]
	tmp := builder.String()
	builder.WriteByte(char)
	lc784Helper(index+1, s, builder, result)
	if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
		char = char ^ (1 << 5)
		builder.Reset()
		builder.WriteString(tmp + string(char))
		lc784Helper(index+1, s, builder, result)
	}
}

// @lc code=end
