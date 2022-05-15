package golang

/*
 * @lc app=leetcode.cn id=812 lang=golang
 *
 * [812] Largest Triangle Area
 */

// @lc code=start
func largestTriangleArea(points [][]int) (ans float64) {
	n := len(points)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				area := lc812Cross(points[j][0]-points[i][0], points[j][1]-points[i][1], points[k][0]-points[i][0], points[k][1]-points[i][1])
				ans = lc812Max(ans, float64(lc812Abs(area))/2.0)
			}
		}
	}
	return ans
}

func lc812Cross(a, b, c, d int) int {
	return a*d - c*b
}

func lc812Max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func lc812Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end
