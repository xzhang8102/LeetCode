package golang

/*
 * @lc app=leetcode.cn id=158 lang=golang
 *
 * [158] Read N Characters Given Read4 II - Call multiple times
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
	tmp := make([]byte, 4)
	ptr := 0   // for tmp
	bound := 0 // for tmp
	return func(buf []byte, n int) int {
		total := 0
		head := 0 // for buf
		for n > 0 {
			size := bound - ptr
			if size > 0 { // content remains in tmp
				if n < size {
					total += n
					copy(buf[head:], tmp[ptr:ptr+n])
					ptr += n
					return total
				} else {
					total += size
					copy(buf[head:], tmp[ptr:bound])
					ptr, bound = 0, 0
					n -= size
					head += size
				}
			} else {
				size = read4(tmp)
				if size < 4 {
					if n < size {
						bound = size
						ptr += n
						total += n
						copy(buf[head:], tmp[:ptr])
					} else {
						total += size
						copy(buf[head:], tmp[:size])
					}
					return total
				} else {
					if n < size {
						bound = size
						ptr += n
						total += n
						copy(buf[head:], tmp[:ptr])
						return total
					} else {
						total += size
						n -= size
						copy(buf[head:], tmp)
						head += size
					}
				}
			}
		}
		return total
	}
}

// @lc code=end
