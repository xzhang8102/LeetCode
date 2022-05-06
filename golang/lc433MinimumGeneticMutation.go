package golang

/*
 * @lc app=leetcode.cn id=433 lang=golang
 *
 * [433] Minimum Genetic Mutation
 */

// @lc code=start
func minMutation(start string, end string, bank []string) int {
	set := map[string]bool{}
	for _, g := range bank {
		set[g] = true
	}
	if !set[end] {
		return -1
	}
	if start == end {
		return 0
	}
	q := []string{start}
	visited := map[string]bool{}
	dist := 0
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, gene := range tmp {
			if gene == end {
				return dist
			}
			visited[gene] = true
			buf := []byte(gene)
			for i := 0; i < 8; i++ {
				g := buf[i]
				for _, c := range "ACGT" {
					buf[i] = byte(c)
					s := string(buf)
					if set[s] && !visited[s] {
						q = append(q, s)
					}
				}
				buf[i] = g
			}
		}
		dist++
	}
	return -1
}

// @lc code=end
