package golang

/*
 * @lc app=leetcode.cn id=1606 lang=golang
 *
 * [1606] Find Servers That Handled Most Number of Requests
 */

// @lc code=start
func busiestServers(k int, arrival []int, load []int) []int {
	type tuple struct{ time, tasks int }
	status := map[int]*tuple{} // time(busy util) - task count
	max := 0
	ans := []int{}
	for i, time := range arrival {
		n := i % k
		cnt := 0
		drop := false
		for t, ok := status[n]; ok && t.time > time; t, ok = status[n] {
			n++
			if n == k {
				n = 0
			}
			cnt++
			if cnt == k {
				drop = true
				break
			}
		}
		if drop {
			continue
		}
		if status[n] == nil {
			status[n] = &tuple{}
		}
		status[n].time = time + load[i]
		status[n].tasks++
		if status[n].tasks == max {
			ans = append(ans, n)
		} else if status[n].tasks > max {
			ans = []int{n}
			max = status[n].tasks
		}
	}
	return ans
}

// @lc code=end
