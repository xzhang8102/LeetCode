package golang

/*
 * @lc app=leetcode.cn id=417 lang=golang
 *
 * [417] Pacific Atlantic Water Flow
 */

// @lc code=start
func pacificAtlantic(heights [][]int) [][]int {
	row := len(heights)
	if row == 0 {
		return nil
	}
	col := len(heights[0])
	pacific, atlantic, visited := make([][]int, row), make([][]int, row), make([][]int, row)
	for i := range visited {
		visited[i] = make([]int, col)
		pacific[i] = make([]int, col)
		atlantic[i] = make([]int, col)
	}
	directions := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	ans := [][]int{}
	var dfs func(int, int, int)
	// flag: 1 - pacific, 2 - atlantic
	dfs = func(r, c, flag int) {
		if flag == 1 {
			if r == 0 || c == 0 {
				pacific[r][c] = 1
			}
		}
		if flag == 2 {
			if r == row-1 || c == col-1 {
				atlantic[r][c] = 1
			}
		}
		visited[r][c] = 1
		for _, dir := range directions {
			newR, newC := r+dir[0], c+dir[1]
			if newR >= 0 && newR < row && newC >= 0 && newC < col && visited[newR][newC] == 0 {
				if flag == 1 && pacific[r][c] == 1 && heights[r][c] <= heights[newR][newC] {
					pacific[newR][newC] = 1
					dfs(newR, newC, flag)
				}
				if flag == 2 && atlantic[r][c] == 1 && heights[r][c] <= heights[newR][newC] {
					atlantic[newR][newC] = 1
					dfs(newR, newC, flag)
				}
			}
		}
		visited[r][c] = 0
	}
	for r := 0; r < row; r++ {
		if r == 0 || r == row-1 {
			for c := 0; c < col; c++ {
				dfs(r, c, 1)
				dfs(r, c, 2)
			}
		} else {
			dfs(r, 0, 1)
			dfs(r, 0, 2)
			dfs(r, col-1, 1)
			dfs(r, col-1, 2)
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if pacific[i][j] == 1 && atlantic[i][j] == 1 {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}

// @lc code=end
