package golang

/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 */

// @lc code=start
type Trie struct {
	isEnd bool
	next  [26]*Trie
}

func Constructor() Trie {
	return Trie{
		false,
		[26]*Trie{},
	}
}

func (this *Trie) Insert(word string) {
	n := len(word)
	for i := 0; i < n; i++ {
		index := word[i] - 'a'
		if node := this.next[index]; node == nil {
			newNode := Constructor()
			this.next[index] = &newNode
			this = &newNode
		} else {
			this = node
		}
	}
	this.isEnd = true
}

func (this *Trie) Search(word string) bool {
	n := len(word)
	for i := 0; i < n; i++ {
		index := word[i] - 'a'
		if node := this.next[index]; node != nil {
			this = node
		} else {
			return false
		}
	}
	return this.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	n := len(prefix)
	for i := 0; i < n; i++ {
		index := prefix[i] - 'a'
		if node := this.next[index]; node != nil {
			this = node
		} else {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
