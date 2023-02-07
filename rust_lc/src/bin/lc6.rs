/*
 * @lc app=leetcode.cn id=6 lang=rust
 *
 * [6] Zigzag Conversion
 */
struct Solution();
// @lc code=start
impl Solution {
    pub fn convert(s: String, num_rows: i32) -> String {
        let n = s.len();
        if num_rows <= 1 || n <= num_rows as usize {
            return s;
        }
        let mut buffer: Vec<Vec<u8>> = vec![Vec::new(); num_rows as usize];
        let mut dir = true;
        let mut line: i32 = 0;
        for (i, c) in s.as_bytes().iter().enumerate() {
            buffer[line as usize].push(*c);
            if line == num_rows - 1 || (line == 0 && i > 0) {
                dir = !dir;
            }
            line += if dir { 1 } else { -1 };
        }
        let mut res = String::new();
        for buf in buffer {
            res.push_str(std::str::from_utf8(&buf).unwrap() );
        }
        res
    }
}
// @lc code=end

fn main() {
    let s = "PAYPALISHIRING".to_string();
    let num_rows = 4;
    assert_eq!("PINALSIGYAHRPI".to_string(), Solution::convert(s, num_rows));
}
