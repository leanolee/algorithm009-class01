//给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。 
//
// 
//
// 示例： 
//二叉树：[3,9,20,null,null,15,7], 
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
// 
//
// 返回其层次遍历结果： 
//
// [
//  [3],
//  [9,20],
//  [15,7]
//]
// 
// Related Topics 树 广度优先搜索
package main

import "fmt"

func main() {
	root := &TreeNode{5, &TreeNode{1, nil, nil}, &TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{6, nil, nil}}}
	order := levelOrder(root)
	fmt.Println(order)
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
func levelOrder(root *TreeNode) [][]int {
	// Stack
	order := make([][]int, 0)
	if root == nil {
		return order
	}
	stack := []*TreeNode{root}
	for n := len(stack); n > 0; n = len(stack) {
		level := make([]int, 0)
		for _, node := range stack {
			level = append(level, node.Val)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		order = append(order, level)
		stack = stack[n:]
	}
	return order

	// 递归
	//order := make([][]int, 0)
	//levelOrderBFS(root, &order, 0)
	//return order
}

func levelOrderBFS(root *TreeNode, order *[][]int, level int) {
	if root == nil {
		return
	}
	if len(*order) == level {
		*order = append(*order, make([]int, 0))
	}
	(*order)[level] = append((*order)[level], root.Val)
	levelOrderBFS(root.Left, order, level+1)
	levelOrderBFS(root.Right, order, level+1)
}

//leetcode submit region end(Prohibit modification and deletion)
