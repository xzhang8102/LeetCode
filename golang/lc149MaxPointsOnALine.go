package golang

/*
 * @lc app=leetcode.cn id=149 lang=golang
 *
 * [149] Max Points on a Line
 */

// @lc code=start
func maxPoints(points [][]int) int {
	n := len(points)
	if n <= 2 {
		return n
	}
	ans := 2
	for p1 := 0; p1 < n; p1++ {
		for p2 := p1 + 1; p2 < n; p2++ {
			count := 2
			for p3 := p2 + 1; p3 < n; p3++ {
				var (
					s1 = (points[p2][1] - points[p1][1]) * (points[p3][0] - points[p2][0])
					s2 = (points[p3][1] - points[p2][1]) * (points[p2][0] - points[p1][0])
				)
				if s1 == s2 {
					count++
				}
			}
			if count > ans {
				ans = count
			}
		}
	}
	return ans
}

// @lc code=end
