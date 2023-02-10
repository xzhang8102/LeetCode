/*
 * @lc app=leetcode.cn id=12 lang=rust
 *
 * [12] Integer to Roman
 */
struct Solution;
// @lc code=start
impl Solution {
    pub fn int_to_roman(num: i32) -> String {
        const S: [(i32, &str); 13]= [
            (1000, "M"),
            (900, "CM"),
            (500, "D"),
            (400, "CD"),
            (100, "C"),
            (90, "XC"),
            (50, "L"),
            (40, "XL"),
            (10, "X"),
            (9, "IX"),
            (5, "V"),
            (4, "IV"),
            (1, "I"),
        ];

        let mut num = num;
        let ret = S.iter().fold(String::new(),|mut n, &r|{
            while num >= r.0{
                n.push_str(r.1);
                num -= r.0;
            }
            n
        });
        ret
    }
}
// @lc code=end
fn main() {
    let num = 1994;
    assert_eq!("MCMXCIV".to_string(), Solution::int_to_roman(num));
}
