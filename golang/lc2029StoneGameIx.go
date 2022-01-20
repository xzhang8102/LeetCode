package golang

/*
 * @lc app=leetcode.cn id=2029 lang=golang
 *
 * [2029] Stone Game IX
 */

// @lc code=start
func stoneGameIX(stones []int) bool {
	stats := [3]int{} // index : stone value % 3
	for _, v := range stones {
		stats[v%3]++
	}
	if stats[0]%2 == 0 {
		return !(stats[1] == 0 || stats[2] == 0)
	}
	return !(lc2029Abs(stats[1]-stats[2]) <= 2)
}

func lc2029Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end
