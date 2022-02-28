package golang

/*
 * @lc app=leetcode.cn id=1601 lang=golang
 *
 * [1601] Maximum Number of Achievable Transfer Requests
 */

// @lc code=start
func maximumRequests(n int, requests [][]int) int {
	ans := 0
	pick := 0
	degrees := make([]int, n)
	var backtrack func(index int)
	backtrack = func(index int) {
		if lc1601Valid(degrees) && pick > ans {
			ans = pick
		}
		for i := index; i < len(requests); i++ {
			r := requests[i]
			degrees[r[0]]--
			degrees[r[1]]++
			pick++
			backtrack(i + 1)
			degrees[r[0]]++
			degrees[r[1]]--
			pick--
		}
	}
	backtrack(0)
	return ans
}

func lc1601Valid(degrees []int) bool {
	for _, v := range degrees {
		if v != 0 {
			return false
		}
	}
	return true
}

// @lc code=end
