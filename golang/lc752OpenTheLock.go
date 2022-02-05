package golang

/*
 * @lc app=leetcode.cn id=752 lang=golang
 *
 * [752] Open the Lock
 */

// @lc code=start
func openLock(deadends []string, target string) int {
	const start = "0000"
	visited := map[string]bool{}
	for _, str := range deadends {
		visited[str] = true
	}
	type pair struct {
		state string
		step  int
	}
	if visited[start] {
		return -1
	}
	q := []pair{{start, 0}}
	visited[start] = true
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if curr.state == target {
			return curr.step
		}
		for _, str := range lc752FindNextStr(curr.state) {
			if !visited[str] {
				visited[str] = true
				q = append(q, pair{str, curr.step + 1})
			}
		}
	}
	return -1
}

func lc752FindNextStr(curr string) (res []string) {
	buf := []byte(curr)
	for i, char := range buf {
		buf[i] = char + 1
		if buf[i] > '9' {
			buf[i] = '0'
		}
		res = append(res, string(buf))
		buf[i] = char - 1
		if buf[i] < '0' {
			buf[i] = '9'
		}
		res = append(res, string(buf))
		buf[i] = char
	}
	return
}

// @lc code=end
