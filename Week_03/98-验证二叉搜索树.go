//给定一个二叉树，判断其是否是一个有效的二叉搜索树。 
//
// 假设一个二叉搜索树具有如下特征： 
//
// 
// 节点的左子树只包含小于当前节点的数。 
// 节点的右子树只包含大于当前节点的数。 
// 所有左子树和右子树自身必须也是二叉搜索树。 
// 
//
// 示例 1: 
//
// 输入:
//    2
//   / \
//  1   3
//输出: true
// 
//
// 示例 2: 
//
// 输入:
//    5
//   / \
//  1   4
//     / \
//    3   6
//输出: false
//解释: 输入为: [5,1,4,null,null,3,6]。
//     根节点的值为 5 ，但是其右子节点值为 4 。
// 
// Related Topics 树 深度优先搜索
package main

import (
	"fmt"
	"math"
)

func main() {
	//root := &TreeNode{5, &TreeNode{1, nil, nil}, &TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{6, nil, nil}}}
	//root := &TreeNode{1, &TreeNode{1, nil, nil}, nil}
	root := &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}
	bst := isValidBST(root)
	fmt.Println(bst)
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
func isValidBST(root *TreeNode) bool {
	// 莫里斯遍历：不推荐，破坏了数据结构
	//curr := root
	//lastVal := math.MinInt64
	//for curr != nil {
	//	if curr.Left == nil {
	//		if curr.Val <= lastVal {
	//			return false
	//		}
	//		lastVal = curr.Val
	//		curr = curr.Right
	//	} else {
	//		pre := curr.Left
	//		for pre.Right != nil {
	//			pre = pre.Right
	//		}
	//		pre.Right, curr, curr.Left = curr, curr.Left, nil
	//	}
	//}
	//return true

	// Stack：中序遍历
	//stack := make([]*TreeNode, 0)
	//curr := root
	//lastVal := math.MinInt64
	//for curr != nil || len(stack) > 0 {
	//	for curr != nil {
	//		stack = append(stack, curr)
	//		curr = curr.Left
	//	}
	//	curr = stack[len(stack)-1]
	//	stack = stack[:len(stack)-1]
	//	fmt.Println(curr.Val, lastVal)
	//	if curr.Val <= lastVal {
	//		return false
	//	}
	//	lastVal = curr.Val
	//	curr = curr.Right
	//}
	//return true

	// dfs：左右边界
	return dfs(math.MinInt64, root, math.MaxInt64)

	// 递归：这样写是通不过的
	//if root == nil {
	//	return true
	//}
	//if root.Left != nil {
	//	if root.Left.Val >= root.Val || !isValidBST(root.Left) {
	//		return false
	//	}
	//}
	//if root.Right != nil {
	//	if root.Right.Val <= root.Val || !isValidBST(root.Right) {
	//		return false
	//	}
	//}
	//return true
}

func dfs(min int, root *TreeNode, max int) bool {
	if root == nil {
		return true
	}
	if min >= root.Val || root.Val >= max {
		return false
	} // 这里没有按中序遍历的顺序
	return dfs(min, root.Left, root.Val) && dfs(root.Val, root.Right, max)
}

//leetcode submit region end(Prohibit modification and deletion)
