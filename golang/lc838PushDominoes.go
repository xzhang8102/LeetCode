package golang

/*
 * @lc app=leetcode.cn id=838 lang=golang
 *
 * [838] Push Dominoes
 */

// @lc code=start
func pushDominoes(dominoes string) string {
	s := []byte(dominoes)
	i, n := 0, len(s)
	leftForce := byte('L')
	for i < n {
		j := i
		for j < n && s[j] == '.' { // 找到一段连续'.'
			j++
		}
		rightForce := byte('R')
		if j < n {
			rightForce = s[j]
		}
		// 检查该段骨牌的两侧的力
		// 如果同向
		if leftForce == rightForce {
			for i < j {
				s[i] = leftForce
				i++
			}
			// 如果相对
		} else if leftForce == byte('R') && rightForce == byte('L') {
			for k := j - 1; i < k; i, k = i+1, k-1 {
				s[i], s[k] = leftForce, rightForce // 两面向中间推
			}
		}
		leftForce = rightForce
		i = j + 1
	}
	return string(s)
}

// @lc code=end
