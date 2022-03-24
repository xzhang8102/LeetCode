package golang

/*
 * @lc app=leetcode.cn id=661 lang=golang
 *
 * [661] Image Smoother
 */

// @lc code=start
func imageSmoother(img [][]int) [][]int {
	row := len(img)
	if row == 0 {
		return nil
	}
	col := len(img[0])
	directions := [9][2]int{
		{0, -1}, {0, 0}, {0, 1},
		{-1, -1}, {-1, 0}, {-1, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	ans := make([][]int, row)
	for i := range ans {
		ans[i] = make([]int, col)
	}
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			sum := 0
			cnt := 0
			for _, dir := range directions {
				if newR, newC := r+dir[0], c+dir[1]; newR >= 0 && newR < row && newC >= 0 && newC < col {
					cnt++
					sum += img[newR][newC]
				}
			}
			ans[r][c] = sum / cnt
		}
	}
	return ans
}

// @lc code=end
