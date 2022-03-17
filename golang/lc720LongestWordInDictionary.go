package golang

/*
 * @lc app=leetcode.cn id=720 lang=golang
 *
 * [720] Longest Word in Dictionary
 */

// @lc code=start
type lc720TridNode struct {
	isEnd bool
	child [26]*lc720TridNode
}

func (trie *lc720TridNode) insert(word string) {
	ptr := trie
	for _, ch := range word {
		idx := ch - 'a'
		if ptr.child[idx] == nil {
			ptr.child[idx] = new(lc720TridNode)
		}
		ptr = ptr.child[idx]
	}
	ptr.isEnd = true
}

func (trie *lc720TridNode) search(word string) bool {
	ptr := trie
	for _, ch := range word {
		idx := ch - 'a'
		if ptr.child[idx] == nil || !ptr.child[idx].isEnd { // every character should be an end
			return false
		}
		ptr = ptr.child[idx]
	}
	return true
}

func longestWord(words []string) string {
	trie := new(lc720TridNode)
	for _, word := range words {
		trie.insert(word)
	}
	longest := ""
	for _, word := range words {
		if trie.search(word) && (len(word) > len(longest) || len(word) == len(longest) && word < longest) {
			longest = word
		}
	}
	return longest
}

// @lc code=end
