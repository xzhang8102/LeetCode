package golang

/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] Word Search II
 */

// @lc code=start
type lc212Trie struct {
	isEnd    bool
	children map[byte]*lc212Trie
}

func findWords(board [][]byte, words []string) []string {
	ans := make([]string, 0)
	if board == nil {
		return ans
	}
	trie := lc212Trie{
		false,
		make(map[byte]*lc212Trie),
	}
	buildTrie(&trie, words)
	row := len(board)
	col := len(board[0])
	var dfs func(r, c int, triePtr *lc212Trie, builder []byte, tmp *[]string, visited [][]bool)
	dfs = func(r, c int, triePtr *lc212Trie, builder []byte, tmp *[]string, visited [][]bool) {
		if r < 0 || r >= row || c < 0 || c >= col {
			return
		}
		if len(triePtr.children) == 0 {
			return
		}
		if visited[r][c] {
			return
		}
		node, ok := triePtr.children[board[r][c]]
		if ok {
			visited[r][c] = true
			builder = append(builder, board[r][c])
			if node.isEnd {
				*tmp = append(*tmp, string(builder))
				node.isEnd = false
			}
			// up
			dfs(r-1, c, node, builder, tmp, visited)
			// down
			dfs(r+1, c, node, builder, tmp, visited)
			// left
			dfs(r, c-1, node, builder, tmp, visited)
			// right
			dfs(r, c+1, node, builder, tmp, visited)
			visited[r][c] = false
		} else {
			return
		}
	}
	// set := make(map[string]bool)
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			tmp := make([]string, 0)
			visited := make([][]bool, row)
			for i := range visited {
				visited[i] = make([]bool, col)
			}
			dfs(r, c, &trie, make([]byte, 0), &tmp, visited)
			// for _, s := range tmp {
			// 	set[s] = true
			// }
			ans = append(ans, tmp...)
		}
	}
	// for k := range set {
	// 	ans = append(ans, k)
	// }
	return ans
}

func buildTrie(trie *lc212Trie, words []string) {
	var ptr *lc212Trie
	for _, word := range words {
		ptr = trie
		for _, char := range []byte(word) {
			node, ok := ptr.children[char]
			if ok {
				ptr = node
			} else {
				newNode := &lc212Trie{
					false,
					map[byte]*lc212Trie{},
				}
				ptr.children[char] = newNode
				ptr = newNode
			}
		}
		ptr.isEnd = true
	}
}

// @lc code=end
