package golang

/*
 * @lc app=leetcode.cn id=467 lang=golang
 *
 * [467] Unique Substrings in Wraparound String
 */

// @lc code=start
func findSubstringInWraproundString(p string) int {
	dp := [26]int{}
	for left, right := 0, 0; right < len(p); {
		if left == right || (p[right]-p[right-1]+26)%26 == 1 {
			dp[p[right]-'a'] = lc467Max(dp[p[right]-'a'], right-left+1)
			right++
		} else {
			left = right
		}
	}
	ans := 0
	for _, v := range dp {
		ans += v
	}
	return ans
}

func lc467Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
