package golang

/*
 * @lc app=leetcode.cn id=673 lang=golang
 *
 * [673] Number of Longest Increasing Subsequence
 */

// @lc code=start
func findNumberOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := 0
	maxLen := 1
	pick := []int{}
	var backtrack func(int)
	backtrack = func(index int) {
		if len(pick) > maxLen {
			maxLen = len(pick)
			ans = 1
		} else if len(pick) == maxLen {
			ans++
		}
		for i := index; i < n; i++ {
			if len(pick) == 0 || nums[i] > pick[len(pick)-1] {
				pick = append(pick, nums[i])
				backtrack(i + 1)
				pick = pick[:len(pick)-1]
			}
		}
	}
	backtrack(0)
	return ans
}

// @lc code=end
