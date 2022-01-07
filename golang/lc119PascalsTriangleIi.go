package golang

/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 */

// @lc code=start
func getRow(rowIndex int) []int {
	cache := [][]int{{1}}
	for i := 1; i <= rowIndex; i++ {
		tmp := make([]int, i+1)
		tmp[0], tmp[i] = 1, 1
		for j := 1; j < i; j++ {
			tmp[j] = cache[i-1][j-1] + cache[i-1][j]
		}
		cache = append(cache, tmp)
	}
	return cache[rowIndex]
}

// @lc code=end
