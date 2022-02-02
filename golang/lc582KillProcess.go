package golang

/*
 * @lc app=leetcode.cn id=582 lang=golang
 *
 * [582] Kill Process
 */

// @lc code=start
func killProcess(pid []int, ppid []int, kill int) []int {
	proc := map[int][]int{}
	for _, id := range pid {
		proc[id] = []int{}
	}
	for i, parent := range ppid {
		if parent != 0 {
			proc[parent] = append(proc[parent], pid[i])
		}
	}
	ans := []int{}
	q := []int{kill}
	for len(q) > 0 {
		ans = append(ans, q[0])
		for _, parent := range proc[q[0]] {
			q = append(q, parent)
		}
		q = q[1:]
	}
	return ans
}

// @lc code=end
