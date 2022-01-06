package golang

/*
 * @lc app=leetcode.cn id=706 lang=golang
 *
 * [706] Design HashMap
 */

// @lc code=start
type mapNode struct {
	key, val int
	next     *mapNode
}

type MyHashMap struct {
	size      int
	container []*mapNode
}

func Constructor() MyHashMap {
	res := MyHashMap{
		1000,
		make([]*mapNode, 1000),
	}
	// 0 <= key, value <= 10^6
	return res
}

func (this *MyHashMap) Put(key int, value int) {
	index := key % this.size
	head := this.container[index]
	if head == nil {
		this.container[index] = &mapNode{key, value, nil}
	} else {
		ptr := head
		for ; ptr.key != key && ptr.next != nil; ptr = ptr.next {
		}
		if ptr.key != key {
			ptr.next = &mapNode{key, value, nil}
		} else {
			ptr.val = value
		}
	}
}

func (this *MyHashMap) Get(key int) int {
	index := key % this.size
	head := this.container[index]
	if head == nil {
		return -1
	} else {
		ptr := head
		for ; ptr.key != key && ptr.next != nil; ptr = ptr.next {
		}
		if ptr.key != key {
			return -1
		} else {
			return ptr.val
		}
	}
}

func (this *MyHashMap) Remove(key int) {
	index := key % this.size
	head := this.container[index]
	if head != nil {
		var prev, curr *mapNode = nil, head
		for ; curr != nil && curr.key != key; prev, curr = curr, curr.next {
		}
		if curr != nil && curr.key == key {
			if prev != nil {
				prev.next = curr.next
			} else {
				this.container[index] = curr.next
			}
			curr.next = nil
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
// @lc code=end
