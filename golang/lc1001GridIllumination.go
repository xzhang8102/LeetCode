package golang

/*
 * @lc app=leetcode.cn id=1001 lang=golang
 *
 * [1001] Grid Illumination
 */

// @lc code=start
func gridIllumination(n int, lamps [][]int, queries [][]int) []int {
	luminDict := map[int]map[int]bool{} // id in grid: id of lamp
	for _, lamp := range lamps {
		for _, id := range lc1001GetSurround(n, n, lamp) {
			lampId := lamp[0]*n + lamp[1]
			if luminDict[id] == nil {
				luminDict[id] = map[int]bool{}
			}
			luminDict[id][lampId] = true
		}
	}
	ans := []int{}
	for _, query := range queries {
		r, c := query[0], query[1]
		if _, ok := luminDict[r*n+c]; ok {
			ans = append(ans, 1)
		} else {
			ans = append(ans, 0)
		}
		for _, off := range lc1001GetSurround(n, 2, query) {
			for id, light := range luminDict {
				if light[off] {
					delete(light, off)
				}
				if len(light) == 0 {
					delete(luminDict, id)
				}
			}
		}
	}
	return ans
}

func lc1001GetSurround(n, limit int, lamp []int) []int {
	directions := [8][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
	}
	res := []int{}
	r, c := lamp[0], lamp[1]
	res = append(res, r*n+c)
	for _, dir := range directions {
		for i := 1; i < n && i < limit; i++ {
			newR, newC := r+i*dir[0], c+i*dir[1]
			if newR < 0 || newR >= n || newC < 0 || newC >= n {
				break
			}
			res = append(res, newR*n+newC)
		}
	}
	return res
}

// @lc code=end
