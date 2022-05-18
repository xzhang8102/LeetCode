package golang

/*
 * @lc app=leetcode.cn id=668 lang=golang
 *
 * [668] Kth Smallest Number in Multiplication Table
 */

// @lc code=start
func findKthNumber(m int, n int, k int) int {
	lo, hi := 1, m*n
	for lo < hi {
		mid := lo + (hi-lo)>>1
		if cnt := lc668GetCnt(m, n, mid); cnt >= k {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return hi
}

func lc668GetCnt(m, n, val int) int {
	ans := 0
	for i := 1; i <= m; i++ {
		if i > val {
			break
		}
		ans += lc668Min(val/i, n)
	}
	return ans
}

func lc668Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
