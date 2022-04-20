package golang

/*
 * @lc app=leetcode.cn id=133 lang=golang
 *
 * [133] Clone Graph
 */

/**
 * Definition for a GraphNode.
 */
type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}

// @lc code=start
func cloneGraph(node *GraphNode) *GraphNode {
	if node == nil {
		return nil
	}
	visited := map[*GraphNode]*GraphNode{node: {Val: node.Val}}
	var dfs func(*GraphNode)
	dfs = func(n *GraphNode) {
		for _, ne := range n.Neighbors {
			if _, seen := visited[ne]; !seen {
				visited[ne] = &GraphNode{Val: ne.Val}
				dfs(ne)
			}
			visited[n].Neighbors = append(visited[n].Neighbors, visited[ne])
		}
	}
	dfs(node)
	return visited[node]
}

// @lc code=end
