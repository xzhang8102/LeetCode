package golang

/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] Course Schedule
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := map[int][]int{}
	visited := make([]int, numCourses)
	valid := true
	for _, pre := range prerequisites {
		to, from := pre[0], pre[1]
		graph[from] = append(graph[from], to)
	}
	var dfs func(start int)
	dfs = func(start int) {
		visited[start] = 1
		for _, next := range graph[start] {
			if visited[next] == 0 {
				dfs(next)
				if !valid {
					return
				}
			} else if visited[next] == 1 {
				valid = false
				return
			}
		}
		visited[start] = 2
	}
	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	return valid
}

// @lc code=end
