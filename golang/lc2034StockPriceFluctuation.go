package golang

import "container/heap"

/*
 * @lc app=leetcode.cn id=2034 lang=golang
 *
 * [2034] Stock Price Fluctuation
 */

// @lc code=start
type pair struct{ price, timestamp int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].price < h[j].price }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

type StockPrice struct {
	maxPrice, minPrice hp
	timePriceMap       map[int]int
	latest             int
}

func lc2034Constructor() StockPrice {
	return StockPrice{timePriceMap: map[int]int{}}
}

func (this *StockPrice) Update(timestamp int, price int) {
	heap.Push(&this.maxPrice, pair{-price, timestamp})
	heap.Push(&this.minPrice, pair{price, timestamp})
	this.timePriceMap[timestamp] = price
	if timestamp > this.latest {
		this.latest = timestamp
	}
}

func (this *StockPrice) Current() int {
	return this.timePriceMap[this.latest]
}

func (this *StockPrice) Maximum() int {
	for {
		if p := this.maxPrice[0]; -p.price == this.timePriceMap[p.timestamp] {
			return -p.price
		}
		heap.Pop(&this.maxPrice)
	}
}

func (this *StockPrice) Minimum() int {
	for {
		if p := this.minPrice[0]; p.price == this.timePriceMap[p.timestamp] {
			return p.price
		}
		heap.Pop(&this.minPrice)
	}
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
// @lc code=end
