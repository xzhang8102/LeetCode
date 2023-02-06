/*
 * @lc app=leetcode.cn id=4 lang=rust
 *
 * [4] Median of Two Sorted Arrays
 */
struct Solution();
// @lc code=start
impl Solution {
    pub fn find_median_sorted_arrays(nums1: Vec<i32>, nums2: Vec<i32>) -> f64 {
        let (len1, len2) = (nums1.len(), nums2.len());
        if len1 == 0 && len2 == 0 {
            return 0.0;
        }
        if len1 > len2 {
            return Self::find_median_sorted_arrays(nums2, nums1);
        }
        let is_even = (len1 + len2) % 2 == 0;
        let half = ((len1 + len2) >> 1) as i32;
        let (mut lo, mut hi) = (0, len1 as i32);
        while lo <= hi { // lo == hi == len1 is allowed, so there is a = sign
            let mid = (lo + hi) >> 1;
            let (mut nums1_l, mut nums2_l) = (std::i32::MIN, std::i32::MIN);
            if mid > 0 {
                nums1_l = nums1[(mid - 1) as usize];
            }
            if half - mid > 0 {
                nums2_l = nums2[(half - mid - 1) as usize];
            }
            let larger = nums1_l.max(nums2_l);
            let (mut nums1_r, mut nums2_r) = (std::i32::MAX, std::i32::MAX);
            if mid < len1 as i32 {
                nums1_r = nums1[mid as usize];
            }
            if half - mid < len2 as i32{
                nums2_r = nums2[(half - mid) as usize];
            }
            let smaller = nums1_r.min(nums2_r);
            if nums1_l <= nums2_r && nums2_l <= nums1_r {
                if is_even {
                    return (larger + smaller) as f64 / 2.0;
                } else {
                    return smaller as f64;
                }
            } else if nums1_l > nums2_r {
                hi = mid - 1;
            } else {
                lo = mid + 1;
            }
        }
        0.0
    }
}
// @lc code=end

fn main() {
    let nums1 = vec![1, 2];
    let nums2 = vec![3, 4];
    assert_eq!(2.5, Solution::find_median_sorted_arrays(nums1, nums2));
}
