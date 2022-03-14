package golang

/*
 * @lc app=leetcode.cn id=798 lang=golang
 *
 * [798] Smallest Rotation with Highest Score
 */

// @lc code=start
// 两点注意：
// 1. 从src index逆时针同向旋转到target index需要(src - target)步
// 2. 差分数组下标为旋转次数，值为差值（类似lc1109）
func bestRotation(nums []int) int {
	n := len(nums)
	diff := make([]int, n+1) // 以差分数组的形式来记录score
	for index, val := range nums {
		if val <= index {
			// rotate left
			diff[0]++
			diff[index-val+1]--
			// rotate right
			diff[index+1]++
			diff[n]--
		} else {
			diff[index+1]++
			diff[n-val+index+1]--
		}
	}
	maxScore := 0
	score := 0
	ans := 0
	for i := 0; i < n; i++ {
		score += diff[i]
		if score > maxScore {
			ans = i
			maxScore = score
		}
	}
	return ans
}

// @lc code=end
