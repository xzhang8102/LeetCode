package golang

/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	if n >= 0 {
		return lc50QuickMul(x, n)
	}
	return 1.0 / lc50QuickMul(x, -n)
}

func lc50QuickMul(x float64, n int) float64 {
	ans := 1.0
	xContribution := x
	for n > 0 {
		if n&1 == 1 {
			ans *= xContribution
		}
		xContribution *= xContribution
		n /= 2
	}
	return ans
}

// @lc code=end
