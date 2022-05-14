package golang

/*
 * @lc app=leetcode.cn id=691 lang=golang
 *
 * [691] Stickers to Spell Word
 */

// @lc code=start
func minStickers(stickers []string, target string) int {
	n := len(target)
	dp := make([]int, 1<<n)
	for i := range dp {
		dp[i] = -1
	}
	var dfs func(state int) int
	dfs = func(state int) int {
		if state == (1<<n)-1 {
			return 0
		}
		if dp[state] != -1 {
			return dp[state]
		}
		dp[state] = n + 1
		for _, sticker := range stickers {
			next := state
		out:
			for _, ch := range sticker {
				for i := 0; i < n; i++ {
					if target[i] == byte(ch) && (next>>i)&1 == 0 {
						next |= 1 << i
						continue out
					}
				}
			}
			if next != state {
				dp[state] = lc691Min(dp[state], dfs(next)+1)
			}
		}
		return dp[state]
	}
	ans := dfs(0)
	if ans == n+1 {
		return -1
	}
	return ans
}

func lc691Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
