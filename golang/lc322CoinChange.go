package golang

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] Coin Change
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	n := len(coins)
	sort.Ints(coins)
	ans := math.MaxInt32
	pick := 0
	var backtrack func(int, int)
	backtrack = func(index, remain int) {
		if remain < 0 {
			return
		}
		if remain == 0 {
			if pick < ans {
				ans = pick
			}
			return
		}
		for i := index; i >= 0; i-- {
			for count := remain / coins[i]; count > 0; count-- {
				pick += count
				if pick < ans {
					backtrack(i-1, remain-count*coins[i])
				}
				pick -= count
			}
		}
	}
	backtrack(n-1, amount)
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

// @lc code=end
