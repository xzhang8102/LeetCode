package golang

/*
 * @lc app=leetcode.cn id=986 lang=golang
 *
 * [986] Interval List Intersections
 */

// @lc code=start
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	ans := [][]int{}
	if len(firstList) == 0 || len(secondList) == 0 {
		return ans
	}
	for i, j := 0, 0; i < len(firstList) && j < len(secondList); {
		if firstList[i][0] > secondList[j][1] {
			j++
		} else if firstList[i][1] < secondList[j][0] {
			i++
		} else {
			left := lc986Max(firstList[i][0], secondList[j][0])
			right := lc986Min(firstList[i][1], secondList[j][1])
			ans = append(ans, []int{left, right})
			if firstList[i][1] > secondList[j][1] {
				j++
			} else if firstList[i][1] < secondList[j][1] {
				i++
			} else {
				i++
				j++
			}
		}
	}
	return ans
}

func lc986Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func lc986Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
