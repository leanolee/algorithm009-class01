//将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。 
//
// 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。 
//
// 示例: 
//
// 给定有序数组: [-10,-3,0,5,9],
//
//一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：
//
//      0
//     / \
//   -3   9
//   /   /
// -10  5
// 
// Related Topics 树 深度优先搜索
package main

import "fmt"

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	bst := sortedArrayToBST(nums)
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
func sortedArrayToBST(nums []int) *TreeNode {
	// 递归
	return sortedArrayToBSTRecursion(nums, 0, len(nums)-1)
}

func sortedArrayToBSTRecursion(arr []int, l int, r int) *TreeNode {
	if l > r {
		return nil
	} else if l == r {
		return &TreeNode{arr[r], nil, nil}
	}
	mid := (l + r + 1) >> 1
	return &TreeNode{arr[mid], sortedArrayToBSTRecursion(arr, l, mid-1), sortedArrayToBSTRecursion(arr, mid+1, r)}
}

//leetcode submit region end(Prohibit modification and deletion)
