package golang

/*
 * @lc app=leetcode.cn id=2039 lang=golang
 *
 * [2039] The Time When the Network Becomes Idle
 */

// @lc code=start
func networkBecomesIdle(edges [][]int, patience []int) int {
	n := len(patience)
	graph := map[int][]int{}
	for _, edge := range edges {
		n1, n2 := edge[0], edge[1]
		graph[n1] = append(graph[n1], n2)
		graph[n2] = append(graph[n2], n1)
	}
	ans := 0
	q := []int{0}
	visited := make([]bool, n)
	visited[0] = true
	for dist := 1; q != nil; dist++ {
		tmp := q
		q = nil
		for _, node := range tmp {
			for _, neighbor := range graph[node] {
				if visited[neighbor] {
					continue
				}
				visited[neighbor] = true
				q = append(q, neighbor)
				ans = lc2039Max(ans, (2*dist-1)/patience[neighbor]*patience[neighbor]+2*dist+1)
			}
		}
	}
	return ans
}

func lc2039Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
