/*
 * @lc app=leetcode.cn id=14 lang=rust
 *
 * [14] Longest Common Prefix
 */
struct Solution;
// @lc code=start
impl Solution {
    pub fn longest_common_prefix(strs: Vec<String>) -> String {
        match strs.is_empty() {
            true => "".to_string(),
            _ => strs.iter().skip(1).fold(strs[0].clone(), |accu, x| {
                accu.chars()
                    .zip(x.chars())
                    .take_while(|(a, b)| a == b)
                    .map(|(a, _)| a)
                    .collect()
            }),
        }
    }
}
// @lc code=end
fn main() {
    let strs = vec![
        "flight".to_string(),
        "flow".to_string(),
        "flower".to_string(),
    ];
    assert_eq!("fl".to_string(), Solution::longest_common_prefix(strs));
}
