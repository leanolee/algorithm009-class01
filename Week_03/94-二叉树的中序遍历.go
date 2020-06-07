//给定一个二叉树，返回它的中序 遍历。 
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
//输出: [1,3,2] 
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？ 
// Related Topics 栈 树 哈希表
package main

import "fmt"

func main() {
	root := &TreeNode{5, &TreeNode{1, nil, nil}, &TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{6, nil, nil}}}
	traversal := inorderTraversal(root)
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
func inorderTraversal(root *TreeNode) []int {
	// 莫里斯遍历
	inorder := make([]int, 0)
	curr := root
	for curr != nil {
		if curr.Left == nil {
			inorder = append(inorder, curr.Val)
			curr = curr.Right
		} else {
			pre := curr.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right, curr, curr.Left = curr, curr.Left, nil
		}
	}
	return inorder

	// Stack
	//inorder := make([]int, 0)
	//stack := make([]*TreeNode, 0)
	//curr := root
	//for curr != nil || len(stack) > 0 {
	//	for curr != nil {
	//		stack = append(stack, curr)
	//		curr = curr.Left
	//	}
	//	curr = stack[len(stack)-1]
	//	stack = stack[:len(stack)-1]
	//	inorder = append(inorder, curr.Val)
	//	curr = curr.Right
	//}
	//return inorder

	// 递归
	//inorder := make([]int, 0)
	//inorderTraversalRecursion(&inorder, root)
	//return inorder
}

func inorderTraversalRecursion(inorder *[]int, root *TreeNode) {
	if root == nil {
		return
	}
	inorderTraversalRecursion(inorder, root.Left)
	*inorder = append(*inorder, root.Val)
	inorderTraversalRecursion(inorder, root.Right)
}

//leetcode submit region end(Prohibit modification and deletion)
