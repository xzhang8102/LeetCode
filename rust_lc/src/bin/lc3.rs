/*
 * @lc app=leetcode.cn id=3 lang=rust
 *
 * [3] Longest Substring Without Repeating Characters
 */
struct Solution();
// @lc code=start
use std::collections::HashMap;

impl Solution {
    pub fn length_of_longest_substring(s: String) -> i32 {
        let n = s.len();
        if n <= 1 {
            return n as i32;
        }
        let (mut left, mut right, mut res) = (0, 0, 0);
        let bytes = s.as_bytes();
        let mut dict: HashMap<u8, usize> = HashMap::new();
        while right < n {
            let c = &bytes[right];
            if let Some(index) = dict.get(c) {
                left = left.max(*index + 1);
            }
            dict.insert(*c, right);
            res = res.max(right - left + 1);
            right += 1;
        }
        res as i32
    }
}
// @lc code=end

fn main() {
    let s = String::from("abba");
    assert_eq!(2, Solution::length_of_longest_substring(s));
}
