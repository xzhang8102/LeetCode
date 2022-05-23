package golang

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=675 lang=golang
 *
 * [675] Cut Off Trees for Golf Event
 */

// @lc code=start
func cutOffTree(forest [][]int) int {
	type pair struct{ dist, r, c int }
	trees := []pair{}
	for i, row := range forest {
		for j, h := range row {
			if h > 1 {
				trees = append(trees, pair{h, i, j})
			}
		}
	}
	sort.Slice(trees, func(i, j int) bool { return trees[i].dist < trees[j].dist })
	n, m := len(forest), len(forest[0])
	directions := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	bfs := func(startR, startC, endR, endC int) int {
		visited := make([][]bool, n)
		for i := range visited {
			visited[i] = make([]bool, m)
		}
		visited[startR][startC] = true
		q := []pair{{0, startR, startC}}
		for len(q) > 0 {
			head := q[0]
			q = q[1:]
			if head.r == endR && head.c == endC {
				return head.dist
			}
			for _, dir := range directions {
				nextR, nextC := head.r+dir[0], head.c+dir[1]
				if nextR >= 0 && nextR < n && nextC >= 0 && nextC < m && !visited[nextR][nextC] && forest[nextR][nextC] >= 1 {
					visited[nextR][nextC] = true
					q = append(q, pair{head.dist + 1, nextR, nextC})
				}
			}
		}
		return -1
	}

	preR, preC := 0, 0
	total := 0
	for _, tree := range trees {
		dist := bfs(preR, preC, tree.r, tree.c)
		if dist == -1 {
			return -1
		}
		total += dist
		preR, preC = tree.r, tree.c
	}
	return total
}

// @lc code=end
