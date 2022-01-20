package golang

/*
 * @lc app=leetcode.cn id=1901 lang=golang
 *
 * [1901] Find a Peak Element II
 */

// @lc code=start
func findPeakGrid(mat [][]int) []int {
	row := len(mat)
	lo, hi := 0, row-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		curr, col := lc1901Max(mat[mid])
		prev, next := -1, -1
		if mid > 0 {
			prev = mat[mid-1][col]
		}
		if mid < row-1 {
			next = mat[mid+1][col]
		}
		if curr > prev && curr > next {
			return []int{mid, col}
		}
		if next > curr {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return nil
}

func lc1901Max(a []int) (int, int) {
	max := a[0]
	maxI := 0
	for i, v := range a[1:] {
		if v > max {
			max = v
			maxI = i + 1
		}
	}
	return max, maxI
}

// @lc code=end
