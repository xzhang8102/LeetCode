package golang

/*
 * @lc app=leetcode.cn id=417 lang=golang
 *
 * [417] Pacific Atlantic Water Flow
 */

// @lc code=start
func pacificAtlantic(heights [][]int) [][]int {
	n, m := len(heights), len(heights[0])
	directions := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	bfs := func(q *[]int, record []bool) {
		for len(*q) > 0 {
			tmp := *q
			*q = []int{}
			for _, coord := range tmp {
				r, c := coord/m, coord%m
				for _, dir := range directions {
					newR, newC := r+dir[0], c+dir[1]
					if newR >= 0 && newR < n && newC >= 0 && newC < m && !record[newR*m+newC] && heights[newR][newC] >= heights[r][c] {
						record[newR*m+newC] = true
						*q = append(*q, newR*m+newC)
					}
				}
			}
		}
	}
	pacific := make([]bool, n*m)
	atlantic := make([]bool, n*m)
	q := []int{}
	for i := 0; i < m; i++ {
		pacific[i] = true
		q = append(q, i)
	}
	for i := 1; i < n; i++ {
		pacific[i*m] = true
		q = append(q, i*m)
	}
	bfs(&q, pacific)
	for i := 0; i < m; i++ {
		atlantic[(n-1)*m+i] = true
		q = append(q, (n-1)*m+i)
	}
	for i := 0; i < n-1; i++ {
		atlantic[i*m+m-1] = true
		q = append(q, i*m+m-1)
	}
	bfs(&q, atlantic)
	ans := [][]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if pacific[i*m+j] && atlantic[i*m+j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}

// @lc code=end
