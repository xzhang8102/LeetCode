package golang

/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] First Missing Positive
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	for i := 0; i < n; i++ {
		num := lc41abs(nums[i])
		// make sure element at the index in the range [1:n] is negative
		if num <= n {
			nums[num-1] = -lc41abs(nums[num-1])
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

func lc41abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end
