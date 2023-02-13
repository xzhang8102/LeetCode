/*
 * @lc app=leetcode.cn id=15 lang=rust
 *
 * [15] 3Sum
 */
struct Solution;
// @lc code=start
impl Solution {
    pub fn three_sum(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let n = nums.len();
        if n < 3 {
            return Vec::new();
        }
        let mut nums = nums;
        nums.sort();
        let mut res: Vec<Vec<i32>> = Vec::new();
        for i in 0..n - 2 {
            if i > 0 && nums[i] == nums[i - 1] {
                continue;
            }
            let (mut j, mut k) = (i + 1, n - 1);
            while j < k {
                let total = nums[i] + nums[j] + nums[k];
                if total == 0 {
                    res.push(vec![nums[i], nums[j], nums[k]]);
                }
                if total <= 0 {
                    j += 1;
                    while j < k && nums[j] == nums[j - 1] {
                        j += 1;
                    }
                }
                if total >= 0 {
                    k -= 1;
                    while j < k && nums[k] == nums[k + 1] {
                        k -= 1;
                    }
                }
            }
        }
        res
    }
}
// @lc code=end
fn main() {
    let nums = vec![1, -1, -1, 0];
    assert_eq!(vec![vec![-1, 0, 1]], Solution::three_sum(nums));
}
