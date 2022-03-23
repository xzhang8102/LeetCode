package golang

/*
 * @lc app=leetcode.cn id=440 lang=golang
 *
 * [440] K-th Smallest in Lexicographical Order
 */

// @lc code=start
type lc440Trie struct {
	val      int
	children [10]*lc440Trie
}

// fill trie based on digit counts
func (trie *lc440Trie) insert(limit int) {
	for i := range trie.children {
		childVal := trie.val*10 + i
		if childVal > 0 && childVal <= limit {
			trie.children[i] = &lc440Trie{val: childVal}
			trie.children[i].insert(limit)
		}
	}
}

func findKthNumber(n int, k int) int {
	trie := &lc440Trie{}
	trie.insert(n)
	ans := 0
	var dfs func(node *lc440Trie)
	dfs = func(node *lc440Trie) {
		if ans != 0 {
			return
		}
		if k == 0 {
			ans = node.val
			return
		}
		for _, child := range node.children {
			if child != nil {
				k--
				dfs(child)
			}
		}
	}
	dfs(trie)
	return ans
}

// @lc code=end
