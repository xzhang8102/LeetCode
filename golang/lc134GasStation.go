package golang

/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] Gas Station
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	start, total, sum := 0, 0, 0
	for i := 0; i < n; i++ {
		diff := gas[i] - cost[i]
		total += diff
		sum += diff
		if sum < 0 {
			sum = 0
			start = i + 1
		}
	}
	if total < 0 {
		return -1
	}
	return start
}

// @lc code=end
