package offer

import "sort"

func findNumberIn2DArray(matrix [][]int, target int) bool {
	n := len(matrix)
	if n == 0 {
		return false
	}
	m := len(matrix[0])
	if m == 0 {
		return false
	}
	for i := 0; i < n; i++ {
		if matrix[i][0] > target {
			break
		} else if matrix[i][m-1] >= target {
			if idx := sort.SearchInts(matrix[i], target); matrix[i][idx] == target {
				return true
			}
		}
	}
	return false
}
