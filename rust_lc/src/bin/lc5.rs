/*
 * @lc app=leetcode.cn id=5 lang=rust
 *
 * [5] Longest Palindromic Substring
 */
struct Solution();
// @lc code=start
impl Solution {
    pub fn longest_palindrome(s: String) -> String {
        let n = s.len();
        if n <= 1 {
            return s;
        }
        let mut dp = vec![false; n];
        dp[n - 1] = true;
        let mut res = &s[n - 1..];
        let mut max_len = 1;
        let s_bytes = s.as_bytes();
        for i in (0..n - 1).rev() {
            let mut tmp = vec![false; n];
            tmp[i] = true;
            for j in i + 1..n {
                let len = j - i + 1;
                tmp[j] = s_bytes[i] == s_bytes[j];
                if len > 2 {
                    tmp[j] = tmp[j] && dp[j - 1];
                }
                if tmp[j] && len > max_len {
                    max_len = len;
                    res = &s[i..j + 1];
                }
            }
            dp = tmp;
        }
        res.to_string()
    }
}
// @lc code=end

fn main() {
    let s = String::from("babad");
    assert_eq!("bab".to_string(), Solution::longest_palindrome(s));
}
