package weekly286

func findDifference(nums1 []int, nums2 []int) [][]int {
	set1, set2 := map[int]bool{}, map[int]bool{}
	ans := make([][]int, 2)
	for _, v := range nums1 {
		set1[v] = true
	}
	for _, v := range nums2 {
		set2[v] = true
	}
	for _, v := range nums1 {
		if _, found := set2[v]; !found && set1[v] {
			ans[0] = append(ans[0], v)
			set1[v] = false
		}
	}
	for _, v := range nums2 {
		if _, found := set1[v]; !found && set2[v] {
			ans[1] = append(ans[1], v)
			set2[v] = false
		}
	}
	return ans
}
