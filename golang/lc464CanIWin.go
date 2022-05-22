package golang

/*
 * @lc app=leetcode.cn id=464 lang=golang
 *
 * [464] Can I Win
 */

// @lc code=start
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if desiredTotal > (1+maxChoosableInteger)*maxChoosableInteger/2 {
		return false
	}
	dp := make([]int8, 1<<maxChoosableInteger)
	for i := range dp {
		dp[i] = -1
	}
	var dfs func(int, int) int8
	dfs = func(usedNum, total int) (res int8) {
		if dp[usedNum] != -1 {
			return dp[usedNum]
		}
		defer func() { dp[usedNum] = res }()
		for i := 0; i < maxChoosableInteger; i++ {
			if usedNum>>i&1 == 0 && (total+i+1 >= desiredTotal || dfs(usedNum|1<<i, total+i+1) == 0) {
				res = 1
				return
			}
		}
		return
	}
	return dfs(0, 0) == 1
}

// @lc code=end
