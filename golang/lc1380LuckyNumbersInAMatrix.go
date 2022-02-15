package golang

/*
 * @lc app=leetcode.cn id=1380 lang=golang
 *
 * [1380] Lucky Numbers in a Matrix
 */

// @lc code=start
func luckyNumbers(matrix [][]int) []int {
	ans := []int{}
	for rowId, row := range matrix {
		colId := lc1380FindMinInRow(row)
		if lc1380IsValidCol(matrix, colId, rowId) {
			ans = append(ans, row[colId])
		}
	}
	return ans
}

func lc1380FindMinInRow(row []int) int {
	col := 0
	for i := 1; i < len(row); i++ {
		if row[i] < row[col] {
			col = i
		}
	}
	return col
}

func lc1380IsValidCol(matrix [][]int, colId, rowId int) bool {
	target := matrix[rowId][colId]
	for i := rowId - 1; i >= 0; i-- {
		if matrix[i][colId] > target {
			return false
		}
	}
	for i := rowId + 1; i < len(matrix); i++ {
		if matrix[i][colId] > target {
			return false
		}
	}
	return true
}

// @lc code=end
