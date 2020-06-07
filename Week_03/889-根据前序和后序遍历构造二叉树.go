//返回与给定的前序和后序遍历匹配的任何二叉树。 
//
// pre 和 post 遍历中的值是不同的正整数。 
//
// 
//
// 示例： 
//
// 输入：pre = [1,2,4,5,3,6,7], post = [4,5,2,6,7,3,1]
//输出：[1,2,3,4,5,6,7]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= pre.length == post.length <= 30 
// pre[] 和 post[] 都是 1, 2, ..., pre.length 的排列 
// 每个输入保证至少有一个答案。如果有多个答案，可以返回其中一个。 
// 
// Related Topics 树
package main

import "fmt"

func main() {
	pre := []int{1, 2, 4, 5, 3, 6, 7}
	post := []int{4, 5, 2, 6, 7, 3, 1}
	prePost := constructFromPrePost(pre, post)
	fmt.Println(prePost)
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
func constructFromPrePost(pre []int, post []int) *TreeNode {
	// 递归：使用索引替代数组
	// 递归
	n := len(pre)
	if n == 0 {
		return nil
	}
	root := &TreeNode{pre[0], nil, nil}
	if n == 1 {
		return root
	}
	lCount := 0 // 左子树长度
	for ; lCount < n; lCount++ {
		if post[lCount] == pre[1] {
			break
		}
	}
	root.Left = constructFromPrePost(pre[1:lCount+2], post[:lCount+1])
	root.Right = constructFromPrePost(pre[lCount+2:], post[lCount+1:n-1])
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
