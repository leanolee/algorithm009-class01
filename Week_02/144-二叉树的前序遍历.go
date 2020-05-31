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
	//root := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
	root := &TreeNode{10, &TreeNode{5, &TreeNode{3, &TreeNode{2, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{7, &TreeNode{6, nil, nil}, &TreeNode{8, nil, nil}}}, &TreeNode{16, &TreeNode{14, nil, nil}, &TreeNode{18, nil, nil}}}
	traversal := preorderTraversal(root)
	fmt.Println(traversal)
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
func preorderTraversal(root *TreeNode) []int {
	// 莫里斯遍历
	preorder := make([]int, 0)
	curr := root
	var pre *TreeNode
	for curr != nil {
		if curr.Left == nil {
			preorder = append(preorder, curr.Val)
			curr = curr.Right
		} else {
			pre = curr.Left
			for pre.Right != nil && pre.Right != curr {
				pre = pre.Right
			}
			if pre.Right == nil {
				preorder = append(preorder, curr.Val)
				pre.Right = curr
				curr = curr.Left
			} else {
				pre.Right = nil
				curr = curr.Right
			}
		}
	}
	return preorder

	// 迭代
	//preorder := make([]int, 0)
	//if root == nil {
	//	return preorder
	//}
	//stack := make([]*TreeNode, 1)
	//stack[0] = root
	//for len(stack) != 0 {
	//	root = stack[len(stack)-1]
	//	stack = stack[:len(stack)-1]
	//	preorder = append(preorder, root.Val)
	//	if root.Right != nil {
	//		stack = append(stack, root.Right)
	//	}
	//	if root.Left != nil {
	//		stack = append(stack, root.Left)
	//	}
	//}
	//return preorder

	// 递归
	//preorder := make([]int, 0)
	//preorderTraversalRecursion(root, &preorder)
	//return preorder
}

func preorderTraversalRecursion(root *TreeNode, preorder *[]int) {
	if root == nil {
		return
	}
	*preorder = append(*preorder, root.Val)
	preorderTraversalRecursion(root.Left, preorder)
	preorderTraversalRecursion(root.Right, preorder)
}

//leetcode submit region end(Prohibit modification and deletion)
