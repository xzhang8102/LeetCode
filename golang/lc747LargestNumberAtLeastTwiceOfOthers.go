package golang

/*
 * @lc app=leetcode.cn id=747 lang=golang
 *
 * [747] Largest Number At Least Twice of Others
 */

// @lc code=start
func dominantIndex(nums []int) int {
	n := len(nums)
	maxN, secondMaxN, maxIndex := -1, -1, 0
	for i := 0; i < n; i++ {
		if nums[i] >= maxN {
			maxIndex = i
			secondMaxN = maxN
			maxN = nums[i]
		} else if nums[i] > secondMaxN {
			secondMaxN = nums[i]
		}
	}
	if maxN >= 2*secondMaxN {
		return maxIndex
	} else {
		return -1
	}
}

// @lc code=end
