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
	f := ans
	for last := n - 1; last >= 0; last-- {
		f += sum - n*nums[last]
		if f > ans {
			ans = f
		}
	}
	return ans
}

// @lc code=end
