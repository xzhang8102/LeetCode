package golang

import "sort"

/*
 * @lc app=leetcode.cn id=1036 lang=golang
 *
 * [1036] Escape a Large Maze
 */

// @lc code=start
func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	if len(blocked) < 2 {
		return true
	}
	sx, sy := source[0], source[1]
	tx, ty := target[0], target[1]
	rows := []int{sx, tx}
	cols := []int{sy, ty}
	for _, pair := range blocked {
		rows = append(rows, pair[0])
		cols = append(cols, pair[1])
	}
	rowMapping, rowLen := lc1036Discrete(rows)
	colMapping, colLen := lc1036Discrete(cols)
	newMap := make([][]bool, rowLen+1)
	for i := range newMap {
		newMap[i] = make([]bool, colLen+1)
	}
	for _, pair := range blocked {
		newR, newC := rowMapping[pair[0]], colMapping[pair[1]]
		newMap[newR][newC] = true
	}
	sx, sy = rowMapping[sx], colMapping[sy]
	tx, ty = rowMapping[tx], colMapping[ty]
	queue := [][]int{{sx, sy}}
	directions := [4][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	newMap[sx][sy] = true
	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		r, c := head[0], head[1]
		if r == tx && c == ty {
			return true
		}
		for _, dir := range directions {
			newR, newC := r+dir[0], c+dir[1]
			if newR >= 0 && newR <= rowLen && newC >= 0 && newC <= colLen &&
				!newMap[newR][newC] {
				queue = append(queue, []int{newR, newC})
				newMap[newR][newC] = true
			}
		}
	}
	return false
}

func lc1036Discrete(lines []int) (mapping map[int]int, newId int) {
	sort.Ints(lines)
	prev := lines[0]
	if prev > 0 {
		newId = 1
	}
	mapping = map[int]int{lines[0]: newId}
	for _, v := range lines[1:] {
		if v-prev == 1 {
			newId++
		} else if v-prev > 1 {
			newId += 2
		}
		mapping[v] = newId
		prev = v
	}
	if lines[len(lines)-1] != 1e6-1 {
		newId++
	}
	return
}

// @lc code=end
