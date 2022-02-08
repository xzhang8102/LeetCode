package golang

/*
 * @lc app=leetcode.cn id=1001 lang=golang
 *
 * [1001] Grid Illumination
 */

// @lc code=start
func gridIllumination(n int, lamps [][]int, queries [][]int) []int {
	// -- row; | col; / left; \ right
	row, col, left, right := map[int]int{}, map[int]int{}, map[int]int{}, map[int]int{}
	lampSet := map[int]bool{}
	for _, lamp := range lamps {
		y, x := lamp[0], lamp[1]
		lampId := y*n + x
		if !lampSet[lampId] {
			lampSet[lampId] = true
			row[y]++
			col[x]++
			left[y-x]++
			right[y+x]++
		}
	}
	ans := make([]int, 0, len(queries))
	directions := [9][2]int{{0, 0}, {0, -1}, {0, 1}, {-1, 0}, {-1, -1}, {-1, 1}, {1, 0}, {1, -1}, {1, 1}}
	for _, query := range queries {
		y, x := query[0], query[1]
		if row[y] > 0 || col[x] > 0 || left[y-x] > 0 || right[y+x] > 0 {
			ans = append(ans, 1)
		} else {
			ans = append(ans, 0)
		}
		for _, dir := range directions {
			newY, newX := y+dir[0], x+dir[1]
			if newY >= 0 && newY < n && newX >= 0 && newX < n && lampSet[newY*n+newX] {
				row[newY]--
				col[newX]--
				left[newY-newX]--
				right[newY+newX]--
				delete(lampSet, newY*n+newX)
			}
		}
	}
	return ans
}

// @lc code=end
