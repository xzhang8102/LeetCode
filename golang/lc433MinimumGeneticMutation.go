package golang

/*
 * @lc app=leetcode.cn id=433 lang=golang
 *
 * [433] Minimum Genetic Mutation
 */

// @lc code=start
func minMutation(start string, end string, bank []string) int {
	graph := map[string][]string{}
	add := func(gene string) {
		buf := []byte(gene)
		for i := 0; i < 8; i++ {
			g := buf[i]
			buf[i] = '*'
			match := string(buf)
			graph[gene] = append(graph[gene], match)
			graph[match] = append(graph[match], gene)
			buf[i] = g
		}
	}
	for _, gene := range bank {
		add(gene)
	}
	if graph[end] == nil {
		return -1
	}
	if start == end {
		return 0
	}
	add(start)
	q := []string{start}
	visited := map[string]bool{start: true}
	dist := 0
	for len(q) > 0 {
		tmp := q
		q = nil
		dist++
		for _, gene := range tmp {
			if gene == end {
				return dist / 2
			}
			for _, next := range graph[gene] {
				if !visited[next] {
					visited[next] = true
					q = append(q, next)
				}
			}
		}
	}
	return -1
}

// @lc code=end
