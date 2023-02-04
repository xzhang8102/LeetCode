/*
 * @lc app=leetcode.cn id=2 lang=rust
 *
 * [2] Add Two Numbers
 */

// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
struct ListNode {
    val: i32,
    next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}

struct Solution();
// @lc code=start
impl Solution {
    pub fn add_two_numbers(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        Solution::helper(l1, l2, 0)
    }

    fn helper(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
        carry: i32,
    ) -> Option<Box<ListNode>> {
        if l1 == None && l2 == None {
            if carry != 0 {
                return Some(Box::new(ListNode::new(carry)));
            }
            return None;
        } else {
            let (v1, n1) = match l1 {
                Some(node) => (node.val, node.next),
                None => (0, None),
            };
            let (v2, n2) = match l2 {
                Some(node) => (node.val, node.next),
                None => (0, None),
            };
            let val = (v1 + v2 + carry) % 10;
            let carry = (v1 + v2 + carry) / 10;
            let mut node = Box::new(ListNode::new(val));
            node.next = Solution::helper(n1, n2, carry);
            Some(node)
        }
    }
}
// @lc code=end

fn main() {
    let l1 = Some(Box::new(ListNode {
        val: 2,
        next: Some(Box::new(ListNode {
            val: 4,
            next: Some(Box::new(ListNode::new(3))),
        })),
    }));
    let l2 = Some(Box::new(ListNode {
        val: 5,
        next: Some(Box::new(ListNode {
            val: 6,
            next: Some(Box::new(ListNode::new(4))),
        })),
    }));
    assert_eq!(
        Some(Box::new(ListNode {
            val: 7,
            next: Some(Box::new(ListNode {
                val: 0,
                next: Some(Box::new(ListNode::new(8)))
            }))
        })),
        Solution::add_two_numbers(l1, l2)
    );
}
