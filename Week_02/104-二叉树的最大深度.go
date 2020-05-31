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

import "fmt"

func main() {
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	depth := maxDepth(root)
	fmt.Println(depth)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
	// Stack
	//level := 0
	//if root == nil {
	//	return level
	//}
	//stack := []*TreeNode{root}
	//for len(stack) != 0 {
	//	level++
	//	n := len(stack)
	//	for _, node := range stack {
	//		if node.Left != nil {
	//			stack = append(stack, node.Left)
	//		}
	//		if node.Right != nil {
	//			stack = append(stack, node.Right)
	//		}
	//	}
	//	stack = stack[n:]
	//}
	//return level

	// 递归
	return maxDepthRecursion(0, 0, root)
}

func maxDepthRecursion(level int, maxL int, root *TreeNode) int {
	if root == nil {
		return level
	}
	left := maxDepthRecursion(level+1, maxL, root.Left)
	right := maxDepthRecursion(level+1, maxL, root.Right)
	if left > right {
		maxL = left
	} else {
		maxL = right
	}
	return maxL
}

//leetcode submit region end(Prohibit modification and deletion)
