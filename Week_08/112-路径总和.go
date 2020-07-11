//给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。 
//
// 说明: 叶子节点是指没有子节点的节点。 
//
// 示例: 
//给定如下二叉树，以及目标和 sum = 22， 
//
//               5
//             / \
//            4   8
//           /   / \
//          11  13  4
//         /  \      \
//        7    2      1
// 
//
// 返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。 
// Related Topics 树 深度优先搜索
package main

import (
	"fmt"
)

func main() {
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	sum := 12
	order := hasPathSum(root, sum)
	fmt.Println(order) // [[3], [9, 20],[15, 7]]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
变种：
	1.任意位置终止
	2.任意位置开始，任意位置终止
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
func hasPathSum(root *TreeNode, sum int) bool {
	// 变种 2：
	return recursion(root, sum, sum)

	// 变种 2：错误
	//if root == nil {
	//	return false
	//}
	//if sum == root.Val {
	//	return true
	//}
	//return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val) || hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)

	// 变种 1
	//if root == nil {
	//	return false
	//}
	//if sum == root.Val {
	//	return true
	//}
	//return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)

	// BFS：
	//if root == nil {
	//	return false
	//}
	//queue := list.New()
	//queue.PushBack(root)
	//memo := list.New()
	//memo.PushBack(root.Val)
	//for queue.Len() > 0 {
	//	curr := queue.Remove(queue.Front()).(*TreeNode)
	//	currVal := memo.Remove(memo.Front()).(int)
	//	if curr.Left == nil && curr.Right == nil {
	//		if currVal == sum {
	//			return true
	//		}
	//		continue
	//	}
	//	if curr.Left != nil {
	//		queue.PushBack(curr.Left)
	//		memo.PushBack(currVal + curr.Left.Val)
	//	}
	//	if curr.Right != nil {
	//		queue.PushBack(curr.Right)
	//		memo.PushBack(currVal + curr.Right.Val)
	//	}
	//}
	//return false

	// DFS
	//if root == nil {
	//	return false
	//}
	//if root.Left == nil && root.Right == nil {
	//	return root.Val == sum
	//}
	//return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
func recursion(root *TreeNode, sum int, curr int) bool {
	if root == nil {
		return false
	}
	if curr == root.Val {
		return true
	}
	return recursion(root.Left, sum, curr-root.Val) || recursion(root.Right, sum, curr-root.Val) || recursion(root.Left, sum, sum) || recursion(root.Right, sum, sum)
}

//leetcode submit region end(Prohibit modification and deletion)
