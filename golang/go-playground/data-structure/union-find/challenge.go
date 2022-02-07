package main

// problem: https://sp19.datastructur.es/materials/clab/clab6/clab6

type bubbleGrid [][]int

func (grid bubbleGrid) popBubbles(darts [][]int) []int {
	row, col := len(grid), len(grid[0])
	dummy := row * col
	uf := NewUnionFind(row*col + 1)
	ans := []int{}
	directions := [4][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	type pair struct {
		r, c int
	}
	for _, dart := range darts {
		dr, dc := dart[0], dart[1]
		if grid[dr][dc] == 1 {
			grid[dr][dc] = 0
			q := []pair{}
			visited := map[int]bool{}
			for i := range grid[0] {
				if grid[0][i] == 1 {
					q = append(q, pair{0, i})
				}
			}
			for len(q) > 0 {
				bubble := q[0]
				q = q[1:]
				id := bubble.r*col + bubble.c
				uf.Union(dummy, id)
				visited[id] = true
				for _, dir := range directions {
					newR, newC := bubble.r+dir[0], bubble.c+dir[1]
					if newR >= 0 && newR < row && newC >= 0 && newC < col && grid[newR][newC] == 1 && !visited[newR*col+newC] {
						q = append(q, pair{newR, newC})
					}
				}
			}
			changed := 0
			for r := 0; r < row; r++ {
				for c := 0; c < col; c++ {
					if grid[r][c] == 1 {
						if !uf.IsConnected(r*col+c, dummy) {
							grid[r][c] = 0
							changed++
						} else {
							uf.Disconnect(r*col + c)
						}
					}
				}
			}
			ans = append(ans, changed)
		} else {
			ans = append(ans, 0)
		}
	}
	return ans
}
