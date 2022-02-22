package golang

import "bytes"

/*
 * @lc app=leetcode.cn id=838 lang=golang
 *
 * [838] Push Dominoes
 */

// @lc code=start
func pushDominoes(dominoes string) string {
	n := len(dominoes)
	q := []int{}
	time := make([]int, n)
	for i := range time {
		time[i] = -1
	}
	force := make([][]byte, n)
	for i, ch := range dominoes {
		if ch != '.' {
			q = append(q, i)
			time[i] = 0
			force[i] = append(force[i], byte(ch))
		}
	}
	ans := bytes.Repeat([]byte{'.'}, n)
	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		if len(force[i]) > 1 { // forces from both sides
			continue
		}
		f := force[i][0]
		ans[i] = f
		next := i - 1
		if f == 'R' {
			next = i + 1
		}
		if next >= 0 && next < n {
			t := time[i]
			if time[next] == -1 {
				q = append(q, next)
				time[next] = t + 1
				force[next] = append(force[next], f)
			} else if time[next] == t+1 {
				force[next] = append(force[next], f)
			}
		}
	}
	return string(ans)
}

// @lc code=end
