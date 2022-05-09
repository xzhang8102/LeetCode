package weekly292

func hasValidPath(grid [][]byte) bool {
	n, m := len(grid), len(grid[0])
	if (m+n-1)&1 == 1 || grid[0][0] == ')' || grid[n-1][m-1] == '(' {
		return false
	}
	visited := make([][][]bool, n)
	for i := range visited {
		visited[i] = make([][]bool, m)
		for j := range visited[i] {
			visited[i][j] = make([]bool, (m+n-1)/2+1) // state value will up to the half of the path
		}
	}
	var dfs func(r, c, state int) bool
	dfs = func(r, c, state int) bool {
		if state > n-r+m-c-1 {
			return false
		}
		if r == n-1 && c == m-1 {
			return state == 1
		}
		if visited[r][c][state] {
			return false
		}
		visited[r][c][state] = true
		if grid[r][c] == '(' {
			state++
		} else if state--; state < 0 {
			return false
		}
		return r < n-1 && dfs(r+1, c, state) || c < m-1 && dfs(r, c+1, state)
	}
	return dfs(0, 0, 0)
}
