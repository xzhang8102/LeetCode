package golang

/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	result := make([]int, 0)
	window := make([]int, 0) // store the index
	push := func(i int) {
		for len(window) > 0 && nums[i] >= nums[window[len(window)-1]] {
			window = window[:len(window)-1]
		}
		window = append(window, i)
	}
	// init window
	for i := 0; i < k; i++ {
		push(i)
	}
	result = append(result, nums[window[0]])
	for i := k; i < n; i++ {
		push(i)
		for window[0] <= i-k {
			window = window[1:]
		}
		result = append(result, nums[window[0]])
	}
	return result
}

// @lc code=end
