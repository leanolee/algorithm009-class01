//给定一个二叉树，返回它的 后序 遍历。 
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
//输出: [3,2,1] 
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？ 
// Related Topics 栈 树
package main

import (
	"./MyTools"
	"fmt"
)

func main() {
	//root := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
	root := &TreeNode{10, &TreeNode{5, &TreeNode{3, &TreeNode{2, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{7, &TreeNode{6, nil, nil}, &TreeNode{8, nil, nil}}}, &TreeNode{16, &TreeNode{14, nil, nil}, &TreeNode{18, nil, nil}}}
	traversal := postorderTraversal(root)
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
func postorderTraversal(root *TreeNode) []int {
	// 莫里斯遍历
	postorder := make([]int, 0)
	curr := root
	for curr != nil {
		if curr.Right == nil {
			postorder = append(postorder, curr.Val)
			curr = curr.Left
		} else {
			pre := curr.Right
			for pre.Left != nil && pre.Left != curr {
				pre = pre.Left
			}
			if pre.Left == nil {
				postorder = append(postorder, curr.Val)
				pre.Left = curr
				curr = curr.Right
			} else {
				pre.Left = nil
				curr = curr.Left
			}
		}
	}
	MyTools.Reverse(postorder)
	return postorder

	// Stack
	//postorder := make([]int, 0)
	//if root == nil {
	//	return postorder
	//}
	//stack := make([]*TreeNode, 1)
	//stack[0] = root
	//var curr *TreeNode
	//for len(stack) != 0 {
	//	curr = stack[len(stack)-1]
	//	stack = stack[:len(stack)-1]
	//	postorder = append(postorder, curr.Val)
	//	if curr.Left != nil {
	//		stack = append(stack, curr.Left)
	//	}
	//	if curr.Right != nil {
	//		stack = append(stack, curr.Right)
	//	}
	//}
	//reverse(postorder)
	//return postorder

	// 递归
	//postorder := make([]int, 0)
	//postorderTraversalRecursion(root, &postorder)
	//return postorder
}

func postorderTraversalRecursion(root *TreeNode, ints *[]int) {
	if root == nil {
		return
	}
	postorderTraversalRecursion(root.Left, ints)
	postorderTraversalRecursion(root.Right, ints)
	*ints = append(*ints, root.Val)
}

//leetcode submit region end(Prohibit modification and deletion)
