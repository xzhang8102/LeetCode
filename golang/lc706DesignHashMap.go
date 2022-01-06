package golang

/*
 * @lc app=leetcode.cn id=706 lang=golang
 *
 * [706] Design HashMap
 */

// @lc code=start
import "container/list"

const base = 769

type entry struct {
	key, val int
}

type MyHashMap struct {
	data []list.List
}

func Constructor() MyHashMap {
	res := MyHashMap{
		make([]list.List, base),
	}
	// 0 <= key, value <= 10^6
	return res
}

func (this *MyHashMap) Put(key int, value int) {
	index := key % base
	for e := this.data[index].Front(); e != nil; e = e.Next() {
		if node := e.Value.(entry); node.key == key {
			e.Value = entry{key, value}
			return
		}
	}
	this.data[index].PushBack(entry{key, value})
}

func (this *MyHashMap) Get(key int) int {
	index := key % base
	for e := this.data[index].Front(); e != nil; e = e.Next() {
		if node := e.Value.(entry); node.key == key {
			return node.val
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	index := key % base
	for e := this.data[index].Front(); e != nil; e = e.Next() {
		if e.Value.(entry).key == key {
			this.data[index].Remove(e)
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
