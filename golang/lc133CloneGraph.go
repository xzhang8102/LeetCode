package golang

/*
 * @lc app=leetcode.cn id=133 lang=golang
 *
 * [133] Clone Graph
 */

/**
 * Definition for a Node.
 */
type Node struct {
	Val       int
	Neighbors []*Node
}

// @lc code=start
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	ans := &Node{Val: node.Val}
	q := []*Node{node}
	visited := map[*Node]*Node{node: ans}
	for len(q) > 0 {
		origin := q[0]
		clone := visited[origin]
		q = q[1:]
		for _, ne := range origin.Neighbors {
			// copy management
			if _, ok := visited[ne]; !ok {
				visited[ne] = &Node{Val: ne.Val}
				q = append(q, ne)
			}
			// keep the shape
			clone.Neighbors = append(clone.Neighbors, visited[ne])
		}
	}
	return ans
}

// @lc code=end
