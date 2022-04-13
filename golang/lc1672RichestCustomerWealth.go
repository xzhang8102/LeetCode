package golang

/*
 * @lc app=leetcode.cn id=1672 lang=golang
 *
 * [1672] Richest Customer Wealth
 */

// @lc code=start
func maximumWealth(accounts [][]int) int {
	max := 0
	for _, customer := range accounts {
		sum := 0
		for _, val := range customer {
			sum += val
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

// @lc code=end
