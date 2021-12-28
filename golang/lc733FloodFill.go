package golang

/*
 * @lc app=leetcode.cn id=733 lang=golang
 *
 * [733] Flood Fill
 */

// @lc code=start
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	var directions = [4][2]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
	row := len(image)
	col := len(image[0])
	var dfs func(r, c, prevColor int)
	dfs = func(r, c, prevColor int) {
		currColor := image[r][c]
		if currColor == '#' || currColor != prevColor {
			return
		}
		image[r][c] = '#'
		for _, dir := range directions {
			newr, newc := r+dir[0], c+dir[1]
			if newr < row && newr >= 0 && newc < col && newc >= 0 {
				dfs(newr, newc, currColor)
			}
		}
		if currColor == prevColor {
			image[r][c] = newColor
		} else {
			image[r][c] = currColor
		}
	}
	dfs(sr, sc, image[sr][sc])
	return image
}

// @lc code=end
