package weekly289

func maxTrailingZeros(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	cache := [1001][2]int{}
	for i := 2; i <= 1000; i++ {
		if i%2 == 0 {
			cache[i][0] = cache[i/2][0] + 1
		}
		if i%5 == 0 {
			cache[i][1] = cache[i/5][1] + 1
		}
	}
	s := make([][][2]int, n)
	for i, row := range grid {
		s[i] = make([][2]int, m+1)
		for j, v := range row {
			s[i][j+1][0] = s[i][j][0] + cache[v][0]
			s[i][j+1][1] = s[i][j][1] + cache[v][1]
		}
	}
	ans := 0
	for c := 0; c < m; c++ {
		cnt2, cnt5 := 0, 0
		// from top to bottom
		for r := 0; r < n; r++ {
			cnt2 += cache[grid[r][c]][0]
			cnt5 += cache[grid[r][c]][1]
			left := s[r][c]
			right := [2]int{s[r][m][0] - s[r][c+1][0], s[r][m][1] - s[r][c+1][1]}
			ans = max(ans, max(min(cnt2+left[0], cnt5+left[1]), min(cnt2+right[0], cnt5+right[1])))
		}
		// from bottom to top
		cnt2, cnt5 = 0, 0
		for r := n - 1; r >= 0; r-- {
			cnt2 += cache[grid[r][c]][0]
			cnt5 += cache[grid[r][c]][1]
			left := s[r][c]
			right := [2]int{s[r][m][0] - s[r][c+1][0], s[r][m][1] - s[r][c+1][1]}
			ans = max(ans, max(min(cnt2+left[0], cnt5+left[1]), min(cnt2+right[0], cnt5+right[1])))
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
