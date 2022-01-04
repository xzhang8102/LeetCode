package golang

/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 */

// @lc code=start

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	ans := 0
	sum := 0
	left, right := 0, 0
	for right < n {
		sum += nums[right]
		for sum >= target {
			if ans == 0 || right-left+1 < ans {
				ans = right - left + 1
			}
			sum -= nums[left]
			left++
		}
		right++
	}
	return ans
}

// @lc code=end
