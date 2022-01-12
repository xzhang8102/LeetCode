package golang

/*
 * @lc app=leetcode.cn id=334 lang=golang
 *
 * [334] Increasing Triplet Subsequence
 */

// @lc code=start
func increasingTriplet(nums []int) bool {
	pick := []int{}
	var backtrack func(int) bool
	backtrack = func(index int) bool {
		if len(pick) == 3 {
			return true
		}
		for i := index; i < len(nums); i++ {
			if len(pick) == 0 || nums[i] > pick[len(pick)-1] {
				pick = append(pick, nums[i])
				if backtrack(i + 1) {
					return true
				}
				pick = pick[:len(pick)-1]
			}
		}
		return false
	}
	return len(nums) > 0 && backtrack(0)
}

// @lc code=end
