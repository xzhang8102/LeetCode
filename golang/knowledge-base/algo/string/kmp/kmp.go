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
	i, j := 0, 1
	for j < n {
		if pattern[i] == pattern[j] {
			table[j] = i + 1 // i - 0 + 1
			i++
		} else {
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
		j++
	}
	return table
}
