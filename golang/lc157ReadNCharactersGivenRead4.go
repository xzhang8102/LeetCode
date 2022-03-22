package golang

/*
 * @lc app=leetcode.cn id=157 lang=golang
 *
 * [157] Read N Characters Given Read4
 */

// @lc code=start
/**
 * The read4 API is already defined for you.
 *
 *     read4 := func(buf4 []byte) int
 *
 * // Below is an example of how the read4 API can be called.
 * file := File("abcdefghijk") // File is "abcdefghijk", initially file pointer (fp) points to 'a'
 * buf4 := make([]byte, 4) // Create buffer with enough space to store characters
 * read4(buf4) // read4 returns 4. Now buf = ['a','b','c','d'], fp points to 'e'
 * read4(buf4) // read4 returns 4. Now buf = ['e','f','g','h'], fp points to 'i'
 * read4(buf4) // read4 returns 3. Now buf = ['i','j','k',...], fp points to end of file
 */

var solution = func(read4 func([]byte) int) func([]byte, int) int {
	// implement read below.
	return func(buf []byte, n int) int {
		total := 0
		ptr := 0
		for n > 0 {
			chars := read4(buf[ptr:])
			if n < chars {
				total += n
				break
			} else {
				total += chars
				if chars < 4 {
					break
				}
				ptr += 4
				n -= 4
			}
		}
		return total
	}
}

// @lc code=end
