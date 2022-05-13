package golang

/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] Course Schedule
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := map[int][]int{}
	inDegree := make([]int, numCourses)
	for _, pre := range prerequisites {
		to, from := pre[0], pre[1]
		graph[from] = append(graph[from], to)
		inDegree[to]++
	}
	q := make([]int, 0, numCourses)
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			q = append(q, i)
		}
	}
	learnt := 0
	for len(q) > 0 {
		course := q[0]
		learnt++
		q = q[1:]
		for _, next := range graph[course] {
			inDegree[next]--
			if inDegree[next] == 0 {
				q = append(q, next)
			}
		}
	}
	return learnt == numCourses
}

// @lc code=end
