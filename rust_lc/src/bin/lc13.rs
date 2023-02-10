/*
 * @lc app=leetcode.cn id=13 lang=rust
 *
 * [13] Roman to Integer
 */
struct Solution;
// @lc code=start
use std::collections::HashMap;

impl Solution {
    pub fn roman_to_int(s: String) -> i32 {
        let map: HashMap<u8, i32> = HashMap::from([
            (b'I', 1),
            (b'V', 5),
            (b'X', 10),
            (b'L', 50),
            (b'C', 100),
            (b'D', 500),
            (b'M', 1000),
        ]);
        let n = s.len();
        let mut res = 0;
        let s_bytes = s.as_bytes();
        for (i, b) in s_bytes.iter().enumerate() {
            let val = map.get(b).unwrap();
            if i < n - 1 && val < map.get(&s_bytes[i+1]).unwrap() {
                res -= val;
            } else {
                res += val;
            }
        }
        res
    }
}
// @lc code=end
fn main() {
    let s = "MCMXCIV".to_string();
    assert_eq!(1994, Solution::roman_to_int(s));
}
