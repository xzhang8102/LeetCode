package golang

/*
 * @lc app=leetcode.cn id=2024 lang=golang
 *
 * [2024] Maximize the Confusion of an Exam
 */

// @lc code=start
func maxConsecutiveAnswers(answerKey string, k int) int {
	return lc2024Max(
		lc2024AssignChar(answerKey, k, 'T'),
		lc2024AssignChar(answerKey, k, 'F'),
	)
}

func lc2024AssignChar(answerKey string, k int, ch byte) int {
	sum := 0
	max := 0
	n := len(answerKey)
	left, right := 0, 0
	for right < n {
		if sum < k || sum == k && answerKey[right] == ch {
			// 将right所指向的字符包含在内，可以避免讨论right==n的情况
			if answerKey[right] != ch {
				sum++
			}
			max = lc2024Max(max, right-left+1)
			right++
		} else {
			if answerKey[left] != ch {
				sum--
			}
			left++
		}
	}
	return max
}

func lc2024Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
