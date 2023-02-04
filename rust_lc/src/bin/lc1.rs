/*
 * @lc app=leetcode.cn id=1 lang=rust
 *
 * [1] Two Sum
 */
struct Solution();
// @lc code=start
use std::collections::HashMap;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut dict: HashMap<i32, usize> = HashMap::new();
        for (i, n) in nums.iter().enumerate() {
            match dict.get(&(&target - n)) {
                Some(index) => return vec![*index as i32, i as i32],
                None => {
                    dict.insert(*n, i);
                }
            }
        }
        Vec::new()
    }
}
// @lc code=end

fn main() {
    let nums = vec![2, 7, 11, 15];
    let target = 9;
    assert_eq!(vec![0, 1], Solution::two_sum(nums, target));
}
