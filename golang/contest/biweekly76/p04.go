package biweekly76

import "sort"

func maximumScore(scores []int, edges [][]int) int {
	n := len(scores)
	graph := make([][]int, n)
	for _, edge := range edges {
		e0, e1 := edge[0], edge[1]
		graph[e0] = append(graph[e0], e1)
		graph[e1] = append(graph[e1], e0)
	}
	for i, conns := range graph {
		if len(conns) > 3 {
			sort.Slice(conns, func(i, j int) bool { return scores[conns[i]] > scores[conns[j]] })
			conns = conns[:3]
			graph[i] = conns
		}
	}
	ans := -1
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		for _, a := range graph[x] {
			if a == y {
				continue
			}
			for _, b := range graph[y] {
				if b != x && a != b {
					ans = max(ans, scores[x]+scores[y]+scores[a]+scores[b])
				}
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
