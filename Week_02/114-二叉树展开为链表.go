//给定一个二叉树，原地将它展开为一个单链表。 
//
// 
//
// 例如，给定二叉树 
//
//     1
//   / \
//  2   5
// / \   \
//3   4   6 
//
// 将其展开为： 
//
// 1
// \
//  2
//   \
//    3
//     \
//      4
//       \
//        5
//         \
//          6 
// Related Topics 树 深度优先搜索
package main

import "fmt"

func main() {
	root := &TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{5, nil, &TreeNode{6, nil, nil}}}
	//var root *TreeNode
	flatten(root)
	for treeNode := root; treeNode != nil; treeNode = treeNode.Right {
		//fmt.Println(treeNode.Left)
		fmt.Println(treeNode.Val)
	}
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
func flatten(root *TreeNode) {
	// morris
	curr := root
	for curr != nil {
		if curr.Left == nil {
			curr = curr.Right
		} else {
			pre := curr.Left
			//for pre.Right != nil && pre.Right != curr {
			for pre.Right != nil {
				pre = pre.Right
			}
			curr.Right, curr.Left, pre.Right = curr.Left, nil, curr.Right
		}
	}

	// Stack
	//if root == nil {
	//	return
	//}
	//stack := []*TreeNode{root}
	//curr := new(TreeNode)
	//for len(stack) != 0 {
	//	next := stack[len(stack)-1]
	//	curr.Right, curr.Left, curr = next, nil, next
	//	stack = stack[:len(stack)-1]
	//	if next.Right != nil {
	//		stack = append(stack, next.Right)
	//	}
	//	if next.Left != nil {
	//		stack = append(stack, next.Left)
	//	}
	//}
	//root = curr.Right
}

//leetcode submit region end(Prohibit modification and deletion)
