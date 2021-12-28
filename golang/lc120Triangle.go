package golang

/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] Triangle
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	row := len(triangle)
	if row <= 1 {
		return triangle[0][0]
	}
	for r := 1; r < row; r++ {
		triangle[r][0] += triangle[r-1][0]
		triangle[r][r] += triangle[r-1][r-1]
		for c := 1; c < r; c++ {
			triangle[r][c] += lc120Min(triangle[r-1][c-1], triangle[r-1][c])
		}
	}
	ans := triangle[row-1][0]
	for _, n := range triangle[row-1] {
		ans = lc120Min(ans, n)
	}
	return ans
}

func lc120Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
