//给定一个二叉树，找出其最大深度。 
//
// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。 
//
// 说明: 叶子节点是指没有子节点的节点。 
//
// 示例： 
//给定二叉树 [3,9,20,null,null,15,7]， 
//
//     3
//   / \
//  9  20
//    /  \
//   15   7 
//
// 返回它的最大深度 3 。 
// Related Topics 树 深度优先搜索
package main

import (
	"./MyTools"
	"fmt"
)

func main() {
	root := &TreeNode{5, &TreeNode{1, nil, nil}, &TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{6, nil, nil}}}
	depth := maxDepth(root)
	fmt.Println(depth)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
递归：BFS
递归：DFS
Stack：BFS
*/
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	// Stack：BFS
	level := 0
	if root == nil {
		return level
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		level++
		n := len(stack)
		for _, node := range stack {
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		}
		stack = stack[n:]
	}
	return level

	// 递归：BFS
	//return bfs(0, root)

	// 递归：DFS
	//if root == nil {
	//	return 0
	//}
	//return MyTools.Max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func bfs(level int, root *TreeNode) int {
	if root == nil {
		return level
	}
	level++
	return MyTools.Max(bfs(level, root.Left), bfs(level, root.Right))
}

//leetcode submit region end(Prohibit modification and deletion)
