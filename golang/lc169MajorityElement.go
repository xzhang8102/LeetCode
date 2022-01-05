package golang

/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start

func majorityElement(nums []int) int {
	candidate := 0
	count := 0
	for _, v := range nums {
		if count == 0 {
			candidate = v // 选举出新的candidate
		}
		if v == candidate {
			count++ // 当前candidate的支持者
		} else {
			count-- // 当前candidate的反对者
		}
	}
	return candidate
}

// @lc code=end
