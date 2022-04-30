package offer

func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, left, right int) int {
	if left >= right {
		return 0
	}
	ans, mid := 0, left+(right-left)>>1
	ans += mergeSort(nums, left, mid)
	ans += mergeSort(nums, mid+1, right)
	tmp := make([]int, right-left+1)
	i, j, k := 0, left, mid+1
	for j <= mid && k <= right {
		if nums[j] <= nums[k] {
			tmp[i] = nums[j]
			j++
		} else {
			tmp[i] = nums[k]
			ans += mid + 1 - j
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
