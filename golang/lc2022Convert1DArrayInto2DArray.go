package golang

/*
 * @lc app=leetcode.cn id=2022 lang=golang
 *
 * [2022] Convert 1D Array Into 2D Array
 */

// @lc code=start
func construct2DArray(original []int, m int, n int) [][]int {
	ans := make([][]int, 0, m)
	L := len(original)
	if L != m*n {
		return ans
	}
	for i := 0; i < L; i = i + n {
		ans = append(ans, original[i:i+n])
	}
	return ans
}

// @lc code=end
