package biweekly77

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}
	for _, g := range guards {
		grid[g[0]][g[1]] = 1
	}
	for _, w := range walls {
		grid[w[0]][w[1]] = 2
	}

	for _, row := range grid {
		for j := 0; j < n; j++ {
			if row[j] == 2 {
				continue
			}
			start, guarded := j, false
			for ; j < n && row[j] != 2; j++ {
				if row[j] == 1 {
					guarded = true
				}
			}
			if guarded {
				for ; start < j; start++ {
					if row[start] == 0 {
						row[start] = -1
					}
				}
			}
		}
	}

	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			if grid[i][j] == 2 {
				continue
			}
			start, guarded := i, false
			for ; i < m && grid[i][j] != 2; i++ {
				if grid[i][j] == 1 {
					guarded = true
				}
			}
			if guarded {
				for ; start < i; start++ {
					if grid[start][j] == 0 {
						grid[start][j] = 1
					}
				}
			}
		}
	}

	ans := 0
	for _, row := range grid {
		for _, v := range row {
			if v == 0 {
				ans++
			}
		}
	}
	return ans
}
