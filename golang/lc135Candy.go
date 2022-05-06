package golang

/*
 * @lc app=leetcode.cn id=135 lang=golang
 *
 * [135] Candy
 */

// @lc code=start
func candy(ratings []int) int {
	n := len(ratings)
	left := make([]int, n)
	left[0] = 1
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	right := 1
	ans := lc135Max(right, left[n-1])
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		ans += lc135Max(right, left[i])
	}
	return ans
}

func lc135Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
