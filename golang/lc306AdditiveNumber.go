package golang

import "strconv"

/*
 * @lc app=leetcode.cn id=306 lang=golang
 *
 * [306] Additive Number
 */

// @lc code=start
func isAdditiveNumber(num string) bool {
	n := len(num)
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			if lc306Valid(i, j, num) {
				return true
			}
		}
	}
	return false
}

func lc306Valid(firstEnd, secondEnd int, num string) bool {
	first, second := num[:firstEnd], num[firstEnd:secondEnd]
	if len(first) > 1 && first[0] == '0' || len(second) > 1 && second[0] == '0' {
		return false
	}
	firstNum, _ := strconv.Atoi(first)
	secondNum, _ := strconv.Atoi(second)
	start := secondEnd
	for start < len(num) {
		sum := strconv.Itoa(firstNum + secondNum)
		if start+len(sum) == len(num) && num[start:] == sum {
			return true
		}
		firstNum, secondNum = secondNum, firstNum+secondNum
		start += len(sum)
	}
	return false
}

// @lc code=end
