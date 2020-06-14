//您需要在二叉树的每一行中找到最大的值。 
//
// 示例： 
//
// 
//输入: 
//
//          1
//         / \
//        3   2
//       / \   \  
//      5   3   9 
//
//输出: [1, 3, 9]
// 
// Related Topics 树 深度优先搜索 广度优先搜索
package main

import "fmt"

func main() {
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	values := largestValues(root)
	fmt.Println(values) // [[3], [9, 20],[15, 7]]
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
func largestValues(root *TreeNode) []int {
	// BFS
	values := make([]int, 0)
	largestValuesBFS(root, &values, 0)
	return values
}

func largestValuesBFS(root *TreeNode, values *[]int, level int) {
	if root == nil {
		return
	}
	if len(*values) == level {
		*values = append(*values, root.Val)
	} else if root.Val > (*values)[level] {
		(*values)[level] = root.Val
	}
	largestValuesBFS(root.Left, values, level+1)
	largestValuesBFS(root.Right, values, level+1)
}

//leetcode submit region end(Prohibit modification and deletion)
