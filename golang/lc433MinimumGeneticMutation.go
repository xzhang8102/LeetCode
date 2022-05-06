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
	qBegin := []string{start}
	qEnd := []string{end}
	seenBegin := map[string]int{start: 0}
	seenEnd := map[string]int{end: 0}
	distBegin := 0
	distEnd := 0
	for len(qBegin) > 0 || len(qEnd) > 0 {
		tmp := qBegin
		qBegin = nil
		for _, gene := range tmp {
			if d, ok := seenEnd[gene]; ok {
				return d + distBegin
			}
			buf := []byte(gene)
			for i := 0; i < 8; i++ {
				g := buf[i]
				for _, c := range "ACGT" {
					buf[i] = byte(c)
					s := string(buf)
					if _, seen := seenBegin[s]; set[s] && !seen {
						seenBegin[s] = distBegin + 1
						qBegin = append(qBegin, s)
					}
				}
				buf[i] = g
			}
		}
		distBegin++
		tmp = qEnd
		qEnd = nil
		for _, gene := range tmp {
			if d, ok := seenBegin[gene]; ok {
				return d + distEnd
			}
			buf := []byte(gene)
			for i := 0; i < 8; i++ {
				g := buf[i]
				for _, c := range "ACGT" {
					buf[i] = byte(c)
					s := string(buf)
					if _, seen := seenEnd[s]; set[s] && !seen {
						seenEnd[s] = distEnd + 1
						qEnd = append(qEnd, s)
					}
				}
				buf[i] = g
			}
		}
		distEnd++
	}
	return -1
}

// @lc code=end
