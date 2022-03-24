package golang

/*
 * @lc app=leetcode.cn id=170 lang=golang
 *
 * [170] Two Sum III - Data structure design
 */

// @lc code=start
type TwoSum struct {
	dict map[int]int
}

func Constructor() TwoSum {
	return TwoSum{map[int]int{}}
}

func (this *TwoSum) Add(number int) {
	this.dict[number]++
}

func (this *TwoSum) Find(value int) bool {
	for v := range this.dict {
		target := value - v
		if target == v && this.dict[v] > 1 || target != v && this.dict[target] >= 1 {
			return true
		}
	}
	return false
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */
// @lc code=end
