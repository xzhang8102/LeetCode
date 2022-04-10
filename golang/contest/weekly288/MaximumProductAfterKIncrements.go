package weekly288

import "sort"

func maximumProduct(nums []int, k int) int {
	const MOD int = 1e9 + 7
	sort.Ints(nums)
	i := 1
	for i < len(nums) && k > 0 {
		if nums[i] == nums[0] {
			i++
		} else {
			gap := nums[i] - nums[0]
			add := min(k/i, gap)
			k -= add * i
			for j := 0; j < i && (add > 0 || k > 0); j++ {
				nums[j] += add
				if add < gap && k > 0 {
					nums[j]++
					k--
				}
			}
		}
	}
	if k > 0 {
		add := k / len(nums)
		k -= add * len(nums)
		for i := range nums {
			nums[i] += add
			if k > 0 {
				nums[i]++
				k--
			}
		}
	}
	ans := 1
	for _, v := range nums {
		ans = ans * v % MOD
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
