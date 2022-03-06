package golang

/*
 * @lc app=leetcode.cn id=2100 lang=golang
 *
 * [2100] Find Good Days to Rob the Bank
 */

// @lc code=start
func goodDaysToRobBank(security []int, time int) []int {
	n := len(security)
	left, right := make([]int, n), make([]int, n)
	for i := 1; i < n; i++ {
		if security[i] <= security[i-1] {
			left[i] = left[i-1] + 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		if security[i+1] >= security[i] {
			right[i] = right[i+1] + 1
		}
	}
	ans := []int{}
	for i := time; i < n-time; i++ {
		if left[i] >= time && right[i] >= time {
			ans = append(ans, i)
		}
	}
	return ans
}

// @lc code=end
