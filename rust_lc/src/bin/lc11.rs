/*
 * @lc app=leetcode.cn id=11 lang=rust
 *
 * [11] Container With Most Water
 */
struct Solution;
// @lc code=start
impl Solution {
    pub fn max_area(height: Vec<i32>) -> i32 {
        let mut res = 0;
        let (mut left, mut right): (i32, i32) = (0, height.len() as i32 - 1);
        while left < right {
            let (left_side, right_side) = (height[left as usize], height[right as usize]);
            res = res.max(left_side.min(right_side) * (right - left));
            if left_side >= right_side {
                right -= 1;
            } else {
                left += 1;
            }
        }
        res
    }
}
// @lc code=end
fn main() {
    let height = vec![1,8,6,2,5,4,8,3,7];
    assert_eq!(49, Solution::max_area(height));
}
