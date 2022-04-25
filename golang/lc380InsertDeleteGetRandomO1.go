package golang

import "math/rand"

/*
 * @lc app=leetcode.cn id=380 lang=golang
 *
 * [380] Insert Delete GetRandom O(1)
 */

// @lc code=start
type RandomizedSet struct {
	list []int
	set  map[int]int
}

func lc380Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.set[val]; ok {
		return false
	}
	this.set[val] = len(this.list)
	this.list = append(this.list, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if idx, ok := this.set[val]; !ok {
		return false
	} else {
		lastVal := this.list[len(this.list)-1]
		this.set[lastVal] = idx
		delete(this.set, val)
		this.list[idx] = lastVal
		this.list = this.list[:len(this.list)-1]
		return true
	}
}

func (this *RandomizedSet) GetRandom() int {
	return this.list[rand.Intn(len(this.list))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end
