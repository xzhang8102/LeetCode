package golang

import "math"

/*
 * @lc app=leetcode.cn id=464 lang=golang
 *
 * [464] Can I Win
 */

// @lc code=start
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if desiredTotal == 0 {
		return true
	}
	if desiredTotal > (1+maxChoosableInteger)*maxChoosableInteger/2 {
		return false
	}
	used := 0
	state := [2][]int{}
	state[0] = make([]int, int(math.Pow(2, float64(maxChoosableInteger))))
	state[1] = make([]int, int(math.Pow(2, float64(maxChoosableInteger))))
	var dfs func(turn, total int) bool
	dfs = func(turn, total int) bool {
		if state[turn&1][used] != 0 {
			return state[turn&1][used] == 1
		}
		if total >= desiredTotal {
			return (turn-1)&1 == 0
		}
		state[turn&1][used] = 2
		if turn&1 == 1 {
			state[1][used] = 1
		}
		for i := 0; i < maxChoosableInteger; i++ {
			if (used>>i)&1 == 0 {
				used ^= 1 << i
				if turn&1 == 0 {
					if dfs(turn+1, total+i+1) {
						used ^= 1 << i
						state[0][used] = 1
						break
					}
				} else {
					if !dfs(turn+1, total+i+1) {
						used ^= 1 << i
						state[1][used] = 2
						break
					}
				}
				used ^= 1 << i
			}
		}
		return state[turn&1][used] == 1
	}
	return dfs(0, 0)
}

// @lc code=end
