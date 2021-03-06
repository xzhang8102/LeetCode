# Notes

## Simple Problems with optimal solutions

1. Find out the nearest smaller / bigger element of the visiting element in the visited part of the array  
Use mono stack (单调栈).  
Example: [lc84](../golang/lc84LargestRectangleInHistogram.go)

2. In-order / post-order tree traversal using BFS (iterative)  
The iterative solution mimics the call stack of the recursive one.  
The key is `stack := []interface{}{}` which saves memory from declaring a `map` of visited nodes.  
Example: [lc606](../golang/lc606ConstructStringFromBinaryTree.go)

## Tricks

1. Find the lowest bit of the given number.  
`i & -i`

2. Calculate the middle element's index during a binary search.  
`mid := lo + (hi - lo) >> 1`

3. When dealing with linked list, there is a great chance that you need a dummy node.

4. Convert the lowest 1 to 0.  
`i & (i-1)`