package main

// space complexity: O(N)
// query time complexity: O(log(N))
// update time complexity: O(log(N))
// creation time complexity: O(Nlog(N))
type BIT struct {
	tree []int
}

// construct the BIT should call update(i+1, num)
// future update should call update(i+1, newVal - oldVal)
func (bit *BIT) Update(index, diff int) {
	for index < len(bit.tree) {
		bit.tree[index] += diff
		index += index & (-index)
	}
}

// query the prefix sum of the original array from index 0 to i
// should call Query(i+1)
func (bit *BIT) Query(index int) int {
	ans := 0
	for index > 0 {
		ans += bit.tree[index]
		index -= index & (-index)
	}
	return ans
}

func NewBIT(nums []int) *BIT {
	bit := &BIT{make([]int, len(nums)+1)}
	for i, num := range nums {
		bit.Update(i+1, num)
	}
	return bit
}
