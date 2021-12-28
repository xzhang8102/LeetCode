package golang

/*
 * @lc app=leetcode.cn id=785 lang=golang
 *
 * [785] Is Graph Bipartite?
 */

// @lc code=start

func isBipartite(graph [][]int) bool {
	var UNCOLORED, RED, GREEN byte = 0, 1, 2
	record := make([]byte, len(graph))
	queue := make([]int, 0)
	color := RED
	for i := range graph {
		if record[i] == UNCOLORED && len(queue) == 0 {
			color = RED
			queue = append(queue, i)
		}
		for len(queue) > 0 {
			size := len(queue)
			for _, v := range queue {
				if record[v] == UNCOLORED {
					record[v] = color
					queue = append(queue, graph[v]...)
				} else if record[v] != color {
					return false
				}
			}
			queue = queue[size:]
			if color == RED {
				color = GREEN
			} else {
				color = RED
			}
		}
	}
	return true
}

// @lc code=end
