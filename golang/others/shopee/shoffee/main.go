package main

import (
	"fmt"
)

// description: https://leetcode.cn/problems/bk9ocA/

func main() {
	var N, K int
	fmt.Scan(&N, &K)
	V := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&V[i])
		V[i] -= K
	}
	fmt.Println(coffeeNerd(V, N, K))
}

func coffeeNerd(V []int, N, K int) int {
	prefixSum := make([]int, N+1)
	for i := 1; i <= N; i++ {
		prefixSum[i] = prefixSum[i-1] + V[i-1]
	}
	return mergeSort(prefixSum, 0, N)
}

// find numbers of non-inverse pair
func mergeSort(nums []int, left, right int) int {
	if left >= right {
		return 0
	}
	ans, mid := 0, left+(right-left)>>1
	ans += mergeSort(nums, left, mid) + mergeSort(nums, mid+1, right)
	tmp := make([]int, right-left+1)
	i, j, k := 0, left, mid+1
	for j <= mid && k <= right {
		if nums[j] <= nums[k] {
			ans += right - k + 1
			tmp[i] = nums[j]
			j++
		} else {
			tmp[i] = nums[k]
			k++
		}
		i++
	}
	if j <= mid {
		copy(tmp[i:], nums[j:mid+1])
	}
	if k <= right {
		copy(tmp[i:], nums[k:right+1])
	}
	copy(nums[left:right+1], tmp)
	return ans
}
