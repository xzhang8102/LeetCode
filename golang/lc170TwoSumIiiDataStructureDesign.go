package golang

/*
 * @lc app=leetcode.cn id=170 lang=golang
 *
 * [170] Two Sum III - Data structure design
 */

// @lc code=start
type TwoSum struct {
	data []int
}

func Constructor() TwoSum {
	return TwoSum{[]int{}}
}

func (this *TwoSum) Add(number int) {
	this.data = append(this.data, number)
}

func (this *TwoSum) Find(value int) bool {
	cache := map[int]bool{}
	for _, v := range this.data {
		if _, ok := cache[value-v]; ok {
			return true
		} else {
			cache[v] = true
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
