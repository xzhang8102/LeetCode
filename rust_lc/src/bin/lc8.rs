/*
 * @lc app=leetcode.cn id=8 lang=rust
 *
 * [8] String to Integer (atoi)
 */
struct Solution();
// @lc code=start
enum State {
    START,
    DIGIT,
    PLUS,
    MINUS,
    END,
}

impl State {
    fn parse_byte(self, byte: &u8) -> Self {
        match byte {
            b' ' => match self {
                Self::START => self,
                _ => Self::END,
            },
            b'0'..=b'9' => Self::DIGIT,
            b'.' | b'a'..=b'z' | b'A'..=b'Z' => Self::END,
            b'+' => match self {
                Self::START => Self::PLUS,
                _ => Self::END,
            },
            b'-' => match self {
                Self::START => Self::MINUS,
                _ => Self::END,
            },
            _ => Self::END,
        }
    }
}

impl Solution {
    pub fn my_atoi(s: String) -> i32 {
        let mut res = 0;
        let mut is_neg = false;
        let mut state = State::START;
        for byte in s.as_bytes() {
            state = state.parse_byte(byte);
            match state {
                State::MINUS => is_neg = true,
                State::DIGIT => {
                    let digit = (*byte - b'0') as i32;
                    if is_neg {
                        if res < i32::MIN / 10 || (res == i32::MIN / 10 && digit > -(i32::MIN % 10)) {
                            return i32::MIN;
                        }
                    } else {
                        if res > i32::MAX / 10 || (res == i32::MAX / 10 && digit > i32::MAX % 10) {
                            return i32::MAX;
                        }
                    }
                    res = res * 10 + if is_neg { -digit } else { digit };
                }
                State::END => return res,
                _ => continue,
            }
        }
        res
    }
}
// @lc code=end
fn main() {
    let s = "    -32 with words".to_string();
    assert_eq!(-32, Solution::my_atoi(s));

    let s1 = "-2147483649".to_string();
    assert_eq!(i32::MIN, Solution::my_atoi(s1));
}
