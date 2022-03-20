package golang

/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {
	ans := 0
	n := len(prices)
	for left, right := 0, 0; right < n; {
		if profit := prices[right] - prices[left]; profit > 0 || right == left {
			right++
			if profit > ans {
				ans = profit
			}
		} else {
			left++
		}
	}
	return ans
}

// @lc code=end
