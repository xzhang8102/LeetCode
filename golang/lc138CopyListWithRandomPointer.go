package golang

/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] Copy List with Random Pointer
 */

/**
 * Definition for a ListNodeWithRandom.
 */
type ListNodeWithRandom struct {
	Val    int
	Next   *ListNodeWithRandom
	Random *ListNodeWithRandom
}

// @lc code=start
func copyRandomList(head *ListNodeWithRandom) *ListNodeWithRandom {
	if head == nil {
		return nil
	}
	for ptr := head; ptr != nil; ptr = ptr.Next.Next {
		next := ptr.Next
		ptr.Next = &ListNodeWithRandom{ptr.Val, next, nil}
	}
	for ptr := head; ptr != nil; ptr = ptr.Next.Next {
		random := ptr.Random
		if random != nil {
			ptr.Next.Random = random.Next
		}
	}
	ans := head.Next
	for ptr1, ptr2 := head, head.Next; ptr1 != nil && ptr2 != nil; ptr1, ptr2 = ptr1.Next, ptr2.Next {
		next := ptr2.Next
		ptr1.Next = next
		if next != nil {
			ptr2.Next = next.Next
		}
	}
	return ans
}

// @lc code=end
