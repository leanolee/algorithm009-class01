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
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	order := levelOrder_(root)
	fmt.Println(order) // [[3], [9, 20],[15, 7]]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*
递归：
BFS：
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
func levelOrder_(root *TreeNode) [][]int {
	// Stack
	order := make([][]int, 0)
	if root == nil {
		return order
	}
	stack := make([]*TreeNode, 1)
	stack[0] = root
	for len(stack) != 0 {
		ord := make([]int, 0)
		n := len(stack)
		for _, node := range stack {
			ord = append(ord, node.Val)
			// 为什么这里要判断 nil，559却不用？429也不用？
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		order = append(order, ord)
		stack = stack[n:]
	}
	return order

	// 递归
	//order := make([][]int, 0)
	//recursionLevelOrder(0, &order, root)
	//return order
}

func recursionLevelOrder(level int, order *[][]int, root *TreeNode) {
	if root == nil {
		return
	}
	if level == len(*order) {
		*order = append(*order, []int{})
	}
	(*order)[level] = append((*order)[level], root.Val)
	recursionLevelOrder(level+1, order, root.Left)
	recursionLevelOrder(level+1, order, root.Right)
}

//leetcode submit region end(Prohibit modification and deletion)
