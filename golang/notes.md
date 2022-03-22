# Notes

## Simple Problems with optimal solutions

1. Find out the first smallest / biggest element of the visiting element in the visited array [#mono stack]()  
Use mono stack (单调栈).  
Example: [lc84](../golang/lc84LargestRectangleInHistogram.go)

2. in-order / post-order tree traversal using BFS (iterative) [#golang specifics]()  
The iterative solution mimics the call stack of the recursive one.  
The key is `stack := []interface{}{}` which saves memory from declaring a `map` of visited nodes.  
Example: [lc606](../golang/lc606ConstructStringFromBinaryTree.go)
