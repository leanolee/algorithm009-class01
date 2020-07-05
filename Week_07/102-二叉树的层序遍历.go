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
	order := levelOrder(root)
	fmt.Println(order) // [[3], [9, 20],[15, 7]]
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
	// dfs
	levelOrder := make([][]int, 0)
	levelOrderDFS(root, &levelOrder, 0)
	return levelOrder
}

func levelOrderDFS(root *TreeNode, ans *[][]int, idx int) {
	if root == nil {
		return
	}
	if len(*ans) == idx {
		*ans = append(*ans, make([]int, 0))
	}
	(*ans)[idx] = append((*ans)[idx], root.Val)
	levelOrderDFS(root.Left, ans, idx+1)
	levelOrderDFS(root.Right, ans, idx+1)
}

//leetcode submit region end(Prohibit modification and deletion)
