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
	//root := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
	root := &TreeNode{10, &TreeNode{5, &TreeNode{3, &TreeNode{2, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{7, &TreeNode{6, nil, nil}, &TreeNode{8, nil, nil}}}, &TreeNode{16, &TreeNode{14, nil, nil}, &TreeNode{18, nil, nil}}}
	traversal := inorderTraversal(root)
	fmt.Println(traversal)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
递归：
迭代：
莫里斯遍历：
	1.curr 记录当前的根节点，pre 记录 curr 的左节点
	2.迭代：如果 curr.left==nil，将当前节点的值保存
	3.如果curr.left!=nil，pre=curr.left，一直遍历pre.right，直到pre.right==nil：
		pre=pre.right
	4.实行交换：
		1.pre.right=curr：左子树的最右底节点的 right = 原根节点
		2.中间变量 temp=curr：
		3.curr=curr.left：原根节点的左节点，变成根节点
		4.temp.left=nil：原根节点，没有了左节点
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
func inorderTraversal(root *TreeNode) []int {
	// 莫里斯遍历
	inorder := make([]int, 0)
	curr := root
	var pre *TreeNode
	for curr != nil {
		if curr.Left == nil {
			inorder = append(inorder, curr.Val)
			curr = curr.Right
		} else {
			pre = curr.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right, curr, curr.Left = curr, curr.Left, nil
		}
	}
	return inorder

	// 迭代：时间复杂度更小？
	//stack := make([]*TreeNode, 0)
	//inorder := make([]int, 0)
	//curr := root
	//for curr != nil || len(stack) != 0 {
	//	// 往回走
	//	if curr == nil {
	//		curr = stack[len(stack)-1]
	//		stack = stack[:len(stack)-1]
	//	} else {
	//		//for root != nil && root.Left != nil {
	//		for curr.Left != nil {
	//			stack = append(stack, curr)
	//			curr = curr.Left
	//		}
	//	}
	//	inorder = append(inorder, curr.Val)
	//	// 一定要有 root == nil 的情况，不然 for root.Left != nil 会陷入死循环
	//	curr = curr.Right
	//}
	//return inorder

	// 迭代
	//stack := make([]*TreeNode, 0)
	//inorder := make([]int, 0)
	//curr := root
	//for curr != nil || len(stack) != 0 {
	//	for curr != nil {
	//		stack = append(stack, curr)
	//		curr = curr.Left
	//	}
	//	// 往回走
	//	curr = stack[len(stack)-1]
	//	stack = stack[:len(stack)-1]
	//	inorder = append(inorder, curr.Val)
	//	curr = curr.Right
	//}
	//return inorder

	// 递归
	//inorder := make([]int, 0)
	//inorderTraversalRecursion(root, &inorder)
	//return inorder
}

func inorderTraversalRecursion(root *TreeNode, inorder *[]int) {
	if root == nil {
		return
	}
	inorderTraversalRecursion(root.Left, inorder)
	*inorder = append(*inorder, root.Val)
	inorderTraversalRecursion(root.Right, inorder)
}

//leetcode submit region end(Prohibit modification and deletion)
