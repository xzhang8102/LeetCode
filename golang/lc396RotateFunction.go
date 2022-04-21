package golang

/*
 * @lc app=leetcode.cn id=396 lang=golang
 *
 * [396] Rotate Function
 */

// @lc code=start
func maxRotateFunction(nums []int) int {
	n := len(nums)
	sum := 0
	ans := 0
	for i, v := range nums {
		ans += v * i
		sum += v
	}
	last := n - 1
	prev := ans
	for i := 1; i < n; i++ {
		curr := (prev - (n-1)*nums[last]) + sum - nums[last]
		if curr > ans {
			ans = curr
		}
		last--
		prev = curr
	}
	return ans
}

// @lc code=end
