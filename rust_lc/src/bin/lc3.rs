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
        let mut left = 0;
        let mut right = 1;
        let mut res = 1;
        let bytes = s.as_bytes();
        let mut dict: HashMap<u8, usize> = HashMap::new();
        dict.insert(bytes[0], 0);
        while right < n {
            let c = &bytes[right];
            match dict.get(c) {
                Some(index) => {
                    let tmp = *index;
                    for i in left..tmp {
                        dict.remove(&bytes[i]);
                    }
                    dict.insert(*c, right);
                    left = tmp + 1;
                }
                None => {
                    dict.insert(*c, right);
                    res = if dict.len() > res { dict.len() } else { res };
                }
            }
            right += 1;
        }
        res as i32
    }
}
// @lc code=end

fn main() {
    let s = String::from("pwwkew");
    assert_eq!(3, Solution::length_of_longest_substring(s));
}
