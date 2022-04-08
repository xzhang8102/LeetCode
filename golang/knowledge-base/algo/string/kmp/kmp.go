package kmp

// build the array of which the ith value is
// the length of the longest suffix which is
// also the prefix of the pattern[:i+1]
//
// Note: the suffix can not be the same as the
// original string
func build(pattern string) []int {
	n := len(pattern)
	table := make([]int, n)
	// i point at the last index of the prefix
	// j point at the character to match
	for i, j := 0, 1; j < n; j++ {
		//    "acacab"
		//      "acacab"
		// -> "acacab"
		//        "acacab"
		// -> "acacab"
		//         "acacab"
		for i > 0 && pattern[i] != pattern[j] {
			i = table[i-1]
		}
		if pattern[i] != pattern[j] {
			table[j] = 0
		} else {
			table[j] = i + 1
			i++
		}
	}
	return table
}

// return the indices of the substring matches the pattern
func KMP(str, pattern string) []int {
	ret := []int{}
	table := build(pattern)
	n := len(str)
	for i, j := 0, 0; i < n; i++ {
		if str[i] == pattern[j] {
			j++
			if j == len(pattern) {
				ret = append(ret, i-len(pattern)+1)
				j = table[j-1]
			}
		} else {
			for j > 0 && str[i] != pattern[j] {
				j = table[j-1]
			}
			if str[i] == pattern[j] {
				j++
			}
		}
	}
	return ret
}
