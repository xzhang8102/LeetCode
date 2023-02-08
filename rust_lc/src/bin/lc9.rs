/*
 * @lc app=leetcode.cn id=9 lang=rust
 *
 * [9] Palindrome Number
 */
struct Solution();
// @lc code=start
impl Solution {
    pub fn is_palindrome(x: i32) -> bool {
        if x < 0 || (x != 0 && x % 10 == 0) {
            return false;
        }
        if x < 10 {
            return true;
        }
        let mut tmp = x;
        let mut n = 0;
        while tmp > 0 {
            tmp /= 10;
            n += 1;
        }
        let is_odd = n & 0x1 == 0x1;
        n = n >> 1;
        let (mut x, mut y) = (x, 0);
        while n > 0 {
            let digit = x % 10;
            y = y * 10 + digit;
            x /= 10;
            n -= 1;
        }
        if is_odd {
            x /= 10;
        }
        x == y
    }
}
// @lc code=end
fn main() {
    let x = 12221;
    assert!(Solution::is_palindrome(x));
}
