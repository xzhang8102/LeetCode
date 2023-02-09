/*
 * @lc app=leetcode.cn id=10 lang=rust
 *
 * [10] Regular Expression Matching
 */
struct Solution;
// @lc code=start
impl Solution {
    pub fn is_match(s: String, p: String) -> bool {
        let (n, m) = (s.len(), p.len());
        let mut dp: Vec<Vec<bool>> = vec![vec![false; m + 1]; n + 1];
        let (s_bytes, p_bytes) = (s.as_bytes(), p.as_bytes());
        dp[0][0] = true;
        // example:
        // s = "ab" p = ".*c"
        // state: (s, p)
        //    ""         .         *          c
        // "" ("", "")T  ("", .)F  ("", .*)T  ("", .*c)F
        // a  (a , "")F  (a , .)T  (a , .*)T  (a , .*c)F
        // b  (ab, "")F  (ab, .)F  (ab, .*)T  (ab, .*c)F
        for row in 0..=n {
            for col in 1..=m {
                let mut char = 0 as u8;
                if row > 0 {
                    char = s_bytes[row-1];
                }
                if p_bytes[col-1] == b'*' {
                    dp[row][col] = dp[row][col - 2];
                    if row > 0 {
                        if p_bytes[col-2] == b'.' || p_bytes[col-2] == char {
                            dp[row][col] = dp[row][col] || dp[row-1][col];
                        }
                    }
                } else {
                    if row > 0 && (p_bytes[col-1] == b'.' || p_bytes[col-1] == char) {
                        dp[row][col] = dp[row-1][col-1];
                    }
                }
            }
        }
        dp[n][m]
    }
}
// @lc code=end
fn main() {
    let (s, p) = ("ab".to_string(), "c*a*b".to_string());
    assert!(Solution::is_match(s, p));

    let (s1, p1) = ("ab".to_string(), ".*c".to_string());
    assert_eq!(false, Solution::is_match(s1, p1));
}
