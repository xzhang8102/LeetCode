package golang

import "container/heap"

/*
 * @lc app=leetcode.cn id=2034 lang=golang
 *
 * [2034] Stock Price Fluctuation
 */

// @lc code=start
type lc2034Tuple struct{ price, timestamp int }
type lc2034Heap []lc2034Tuple

func (h lc2034Heap) Len() int            { return len(h) }
func (h lc2034Heap) Less(i, j int) bool  { return h[i].price < h[j].price }
func (h lc2034Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *lc2034Heap) Push(v interface{}) { *h = append(*h, v.(lc2034Tuple)) }
func (h *lc2034Heap) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

type StockPrice struct {
	maxPrice, minPrice lc2034Heap
	timePriceMap       map[int]int
	latest             int
}

func lc2034Constructor() StockPrice {
	return StockPrice{timePriceMap: map[int]int{}}
}

func (this *StockPrice) Update(timestamp int, price int) {
	heap.Push(&this.maxPrice, lc2034Tuple{-price, timestamp})
	heap.Push(&this.minPrice, lc2034Tuple{price, timestamp})
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
