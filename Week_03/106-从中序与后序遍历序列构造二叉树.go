//根据一棵树的中序遍历与后序遍历构造二叉树。 
//
// 注意: 
//你可以假设树中没有重复的元素。 
//
// 例如，给出 
//
// 中序遍历 inorder = [9,3,15,20,7]
//后序遍历 postorder = [9,15,7,20,3] 
//
// 返回如下的二叉树： 
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
// 
// Related Topics 树 深度优先搜索 数组
package main

import (
	"fmt"
)

func main() {
	postorder := []int{9, 15, 7, 20, 3}
	inorder := []int{9, 3, 15, 20, 7}
	//preorder := []int{1, 2}
	//inorder := []int{2, 1}
	tree := buildTree(inorder, postorder)
	fmt.Println(tree)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
迭代：
	参考：leetcode-105：前序 + 中序
		前序：根左右
		中序：左根右
	后序：左右根
		反转后：	根右左
	中序：左根右	右根左
		即为leetcode-105的 倒序遍历 + 左右反转
递归：
	写法一
	写法二
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	// 迭代
	n := len(postorder)
	if n == 0 {
		return nil
	}
	root := &TreeNode{postorder[n-1], nil, nil}
	stack := []*TreeNode{root}
	idx := len(inorder) - 1
	for i := n - 2; i >= 0; i-- {	// 注意是 -2
		currVal := postorder[i]
		pre := stack[len(stack)-1]
		if pre.Val != inorder[idx] {
			pre.Right = &TreeNode{currVal, nil, nil}
			stack = append(stack, pre.Right)
		} else {
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[idx] {
				pre = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				idx--
			}
			pre.Left = &TreeNode{currVal, nil, nil}
			stack = append(stack, pre.Left)
		}
	}
	return root

	// 递归
	//memo := make(map[int]int)
	//for i, val := range inorder {
	//	memo[val] = i
	//}
	////return buildTreeDFS_(inorder, postorder, memo, 0, len(postorder)-1, 0, len(inorder)-1)
	//last = len(postorder) - 1
	//return buildTreeDFS(inorder, postorder, memo, 0, len(inorder)-1, len(postorder)-1)
}

var last int

// 写法二 不好理解
func buildTreeDFS(inorder []int, postorder []int, memo map[int]int, inLeft int, inRight int, lastPost int) *TreeNode {
	if inLeft > inRight {
		return nil
	}
	inIdx := memo[postorder[last]]
	root := &TreeNode{postorder[last], nil, nil}
	last-- // 后序的倒序是右子树优先遍历，last-- 正好对应相应的 left 节点
	root.Right = buildTreeDFS(inorder, postorder, memo, inIdx+1, inRight, lastPost-1)
	root.Left = buildTreeDFS(inorder, postorder, memo, inLeft, inIdx-1, lastPost-1)
	return root
}

// 写法一 和leetcode-105相同模板
func buildTreeDFS_(inorder []int, postorder []int, memo map[int]int, pLeft int, pRight int, inLeft int, inRight int) *TreeNode {
	if pLeft > pRight {
		return nil
	}
	inIdx := memo[postorder[pRight]]
	curr := &TreeNode{postorder[pRight], nil, nil}
	curr.Left = buildTreeDFS_(inorder, postorder, memo, pLeft, inIdx-inLeft+pLeft-1, inLeft, inIdx-1)
	curr.Right = buildTreeDFS_(inorder, postorder, memo, inIdx-inLeft+pLeft, pRight-1, inIdx+1, inRight)
	return curr
}

//leetcode submit region end(Prohibit modification and deletion)
