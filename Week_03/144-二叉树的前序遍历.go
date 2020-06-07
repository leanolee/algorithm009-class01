//给定一个二叉树，返回它的 前序 遍历。 
//
// 示例: 
//
// 输入: [1,null,2,3]  
//   1
//    \
//     2
//    /
//   3 
//
//输出: [1,2,3]
// 
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？ 
// Related Topics 栈 树
package main

import "fmt"

func main() {
	root := &TreeNode{5, &TreeNode{1, nil, nil}, &TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{6, nil, nil}}}
	traversal := preorderTraversal(root)
	fmt.Println(traversal)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*

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
func preorderTraversal(root *TreeNode) []int {
	// 莫里斯
	//preorder := make([]int, 0)
	//curr := root
	//for curr != nil {
	//	if curr.Left == nil {
	//		preorder = append(preorder, curr.Val)
	//		curr = curr.Right
	//	} else {
	//		pre := curr.Left
	//		for pre.Right != nil && pre.Right != curr {
	//			pre = pre.Right
	//		}
	//		if pre.Right == nil {
	//			preorder = append(preorder, curr.Val)
	//			pre.Right, curr = curr, curr.Left
	//		} else {
	//			pre.Right, curr = nil, curr.Right
	//		}
	//	}
	//}
	//return preorder

	// Stack
	preorder := make([]int, 0)
	if root == nil {
		return preorder
	}
	stack := []*TreeNode{root}
	var curr *TreeNode
	for len(stack) > 0 {
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		preorder = append(preorder, curr.Val)
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
	}
	return preorder

	// 递归
	//preorder := make([]int, 0)
	//preorderTraversalDFS(root, &preorder)
	//return preorder
}

func preorderTraversalDFS(root *TreeNode, preorder *[]int) {
	if root == nil {
		return
	}
	*preorder = append(*preorder, root.Val)
	preorderTraversalDFS(root.Left, preorder)
	preorderTraversalDFS(root.Right, preorder)
}

//leetcode submit region end(Prohibit modification and deletion)
