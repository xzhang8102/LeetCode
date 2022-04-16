package biweekly76

func maximumScore(scores []int, edges [][]int) int {
	n := len(scores)
	graph := make([][]int, n)
	for _, edge := range edges {
		e0, e1 := edge[0], edge[1]
		graph[e0] = append(graph[e0], e1)
		graph[e1] = append(graph[e1], e0)
	}
	ans := -1
	visited := map[int]bool{}
	picked := []int{}
	var dfs func(start, total int)
	dfs = func(start, total int) {
		if len(picked) == 4 {
			ans = total
			return
		}
		if start == -1 {
			for i := 0; i < n; i++ {
				visited[i] = true
				picked = append(picked, i)
				for _, next := range graph[i] {
					visited[next] = true
					picked = append(picked, next)
					dfs(next, scores[i]+scores[next])
					visited[next] = false
					picked = picked[:len(picked)-1]
				}
				visited[i] = false
				picked = picked[:len(picked)-1]
			}
		} else {
			for _, next := range graph[start] {
				if !visited[next] && (len(picked) < 3 || (len(picked) == 3 && total+scores[next] > ans)) {
					visited[next] = true
					picked = append(picked, next)
					dfs(next, total+scores[next])
					visited[next] = false
					picked = picked[:len(picked)-1]
				}
			}
		}
	}
	dfs(-1, 0)
	return ans
}
