/*
 * @lc app=leetcode.cn id=7 lang=rust
 *
 * [7] Reverse Integer
 */
struct Solution();
// @lc code=start
impl Solution {
    pub fn reverse(x: i32) -> i32 {
        let mut target = x;
        let mut res: i32 = 0;
        let (lo, hi) = (i32::MIN / 10, i32::MAX / 10);
        while target != 0 {
            if res < lo || res > hi {
                return 0;
            }
            let next = target % 10;
            res = res * 10 +next;
            target /= 10;
        }
        res
    }
}
// @lc code=end

fn main() {
    let x = -123;
    assert_eq!(-321, Solution::reverse(x));

    let y = std::i32::MIN;
    assert_eq!(0, Solution::reverse(y));
}
