package offerii

import "strings"

// https://leetcode.cn/problems/Jf1JuT/

func alienOrder(words []string) string {
	graph := map[byte][]byte{}
	inDeg := map[byte]int{}
	for _, c := range words[0] {
		inDeg[byte(c)] = 0
	}
	for i := 1; i < len(words); i++ {
		s, t := words[i-1], words[i]
		for _, c := range t {
			inDeg[byte(c)] += 0
		}
		j := 0
		for j < len(s) && j < len(t) && s[j] == t[j] {
			j++
		}
		if j == len(s) {
			continue
		}
		if j == len(t) && j < len(s) {
			return ""
		}
		graph[s[j]] = append(graph[s[j]], t[j])
		inDeg[t[j]]++
	}
	q := []byte{}
	var builder strings.Builder
	for c, d := range inDeg {
		if d == 0 {
			q = append(q, c)
			builder.WriteByte(c)
		}
	}
	for len(q) > 0 {
		head := q[0]
		q = q[1:]
		for _, next := range graph[head] {
			inDeg[next]--
			if inDeg[next] == 0 {
				q = append(q, next)
				builder.WriteByte(next)
			}
		}
	}
	if builder.Len() == len(inDeg) {
		return builder.String()
	}
	return ""
}
