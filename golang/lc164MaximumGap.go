package golang

/*
 * @lc app=leetcode.cn id=164 lang=golang
 *
 * [164] Maximum Gap
 */

// @lc code=start
func maximumGap(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	buf := make([]int, n)
	maxVal := lc164Max(nums...)
	for exp := 1; exp <= maxVal; exp *= 10 {
		cnt := [10]int{}
		for _, v := range nums {
			digit := v / exp % 10
			cnt[digit]++
		}
		for i := 1; i < 10; i++ {
			cnt[i] += cnt[i-1]
		}
		for i := n - 1; i >= 0; i-- {
			digit := nums[i] / exp % 10
			buf[cnt[digit]-1] = nums[i]
			cnt[digit]--
		}
		copy(nums, buf)
	}
	ans := 0
	for i := 1; i < n; i++ {
		ans = lc164Max(ans, nums[i]-nums[i-1])
	}
	return ans
}

func lc164Max(nums ...int) int {
	ret := nums[0]
	for _, v := range nums[1:] {
		if v > ret {
			ret = v
		}
	}
	return ret
}

// @lc code=end
