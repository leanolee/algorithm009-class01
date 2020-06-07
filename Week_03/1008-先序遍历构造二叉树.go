//返回与给定先序遍历 preorder 相匹配的二叉搜索树（binary search tree）的根结点。 
//
// (回想一下，二叉搜索树是二叉树的一种，其每个节点都满足以下规则，对于 node.left 的任何后代，值总 < node.val，而 node.right
// 的任何后代，值总 > node.val。此外，先序遍历首先显示节点的值，然后遍历 node.left，接着遍历 node.right。） 
//
// 
//
// 示例： 
//
// 输入：[8,5,1,7,10,12]
//输出：[8,5,10,1,7,null,12]
//
// 
//
// 
//
// 提示： 
//
// 
// 1 <= preorder.length <= 100 
// 先序 preorder 中的值是不同的。 
// 
// Related Topics 树
package main

import "fmt"

func main() {
	preorder := []int{8, 5, 1, 7, 10, 12}
	fromPreorder := bstFromPreorder(preorder)
	fmt.Println(fromPreorder)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
递归：先+中
	先排序，得到先序遍历 inorder 数组
	再结合，leetcode-105 获得二叉树
递归：获取左右边界
递归：lower + upper
	参见：https://leetcode-cn.com/problems/construct-binary-search-tree-from-preorder-traversal/solution/jian-kong-er-cha-shu-by-leetcode/
Stack：lower + upper 的变形
	参见：https://leetcode-cn.com/problems/construct-binary-search-tree-from-preorder-traversal/solution/jian-kong-er-cha-shu-by-leetcode/
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
func bstFromPreorder(preorder []int) *TreeNode {
	// Stack：lower + upper 的变形


	// 递归：lower + upper
	// 递归：获取左右边界
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	if len(preorder) == 1 {
		return root
	}
	rIdx := 1 // 右子树起始位置
	for ; rIdx < len(preorder); rIdx++ {
		if preorder[rIdx] > preorder[0] {
			break
		}
	}
	root.Left = bstFromPreorder(preorder[1:rIdx])
	root.Right = bstFromPreorder(preorder[rIdx:])
	return root
	// 递归：先+中
}

//leetcode submit region end(Prohibit modification and deletion)
