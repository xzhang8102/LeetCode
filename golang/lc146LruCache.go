package golang

import "container/list"

/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU Cache
 */

// @lc code=start

type LRUCache struct {
	*list.List
	data     map[int]*list.Element
	capacity int
}

func lc146Constructor(capacity int) LRUCache {
	return LRUCache{
		list.New(),
		map[int]*list.Element{},
		capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.data[key]; ok {
		this.List.MoveToBack(node)
		return node.Value.([]int)[1]
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.data[key]; ok {
		node.Value.([]int)[1] = value
		this.List.MoveToBack(node)
	} else {
		if len(this.data) == this.capacity {
			lru := this.List.Front()
			delete(this.data, lru.Value.([]int)[0])
			this.List.Remove(lru)
		}
		this.data[key] = this.List.PushBack([]int{key, value})
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
