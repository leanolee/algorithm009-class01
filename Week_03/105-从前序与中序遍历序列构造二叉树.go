//根据一棵树的前序遍历与中序遍历构造二叉树。 
//
// 注意: 
//你可以假设树中没有重复的元素。 
//
// 例如，给出 
//
// 前序遍历 preorder = [3,9,20,15,7]
//中序遍历 inorder = [9,3,15,20,7] 
//
// 返回如下的二叉树： 
//
//     3
//   / \
//  9  20
//    /  \
//   15   7 
// Related Topics 树 深度优先搜索 数组
package main

import "fmt"

func main() {
	//preorder := []int{3, 9, 20, 15, 7}
	//inorder := []int{9, 3, 15, 20, 7}
	preorder := []int{8, 5, 1, 7, 10, 12}
	inorder := []int{1, 5, 7, 8, 10, 12}
	//preorder := []int{1, 2}
	//inorder := []int{2, 1}
	tree := buildTree(preorder, inorder)
	fmt.Println(tree)
	values := getValues(tree)
	fmt.Println(values)
}

func getValues(tree *TreeNode) [][]int {
	// 层序遍历，测试结果
	values := make([][]int, 0)
	stack := []*TreeNode{tree}
	level := 0
	for ; len(stack) > 0; level++ {
		if len(values) == level {
			values = append(values, make([]int, 0))
		}
		n := len(stack)
		for _, node := range stack {
			if node == nil {
				values[level] = append(values[level], 0)
			}
			if node != nil {
				values[level] = append(values[level], node.Val)
				stack = append(stack, node.Left)
				stack = append(stack, node.Right)
			}
		}
		stack = stack[n:]
	}
	return values
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
递归：
	前序遍历数组：preorder
	中序遍历数组：inorder
	1.将 inorder 映射为 memo：key: inorder[i], value: i
	2.构建递归函数：buildTreeDFS(preorder, inorder, memo, preLeft, preRight, inLeft, inRight)
		preLeft：当前子树的起始位置
		preRight：当前子树的终止位置
	3.分治：
		curr：当前子树的根节点，value 为前序的第一个元素：preorder[preLeft]
		inIdx：curr 在 inorder 中的位置：memo[preorder[preLeft]]
		下一次左子树范围：
			preLeft：preLeft + 1
			preRight：inIdx - inLeft + preLeft
			inLeft：inLeft
			inRight：inIdx - 1
		下一次右子树范围：
			preLeft：inIdx - inLeft + preLeft + 1
			preRight：preRight
			inLeft：inIdx + 1
			inRight：inRight
迭代：
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 迭代
	n := len(preorder)
	if n == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	stack := []*TreeNode{root}
	idx := 0
	for i := 1; i < n; i++ { // 注意是 1
		currVal := preorder[i]
		pre := stack[len(stack)-1]
		if pre.Val != inorder[idx] { // 左
			pre.Left = &TreeNode{currVal, nil, nil}
			stack = append(stack, pre.Left)
		} else { // 右
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[idx] {
				pre = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				idx++ // 计算有右节点的节点的位置
			}
			pre.Right = &TreeNode{currVal, nil, nil}
			stack = append(stack, pre.Right)
		}
	}
	return root

	// 递归
	//memo := make(map[int]int)
	//for i, val := range inorder { // k: nodeVal, v: index
	//	memo[val] = i
	//}
	//return buildTreeDFS_(preorder, inorder, memo, 0, len(preorder)-1, 0, len(inorder)-1)
	//return buildTreeDFS(preorder, memo, 0, len(preorder)-1, 0)
}

// 优化
func buildTreeDFS(preorder []int, memo map[int]int, preLeft int, preRight int, inLeft int) *TreeNode {
	if preLeft > preRight { // 确实存在这种情况
		return nil
	}
	if preLeft == preRight { // 有了上面判断，这个判断可要可不要。有这判断，更快
		return &TreeNode{preorder[preLeft], nil, nil}
	}
	inIdx := memo[preorder[preLeft]]
	return &TreeNode{preorder[preLeft], buildTreeDFS(preorder, memo, preLeft+1, inIdx-inLeft+preLeft, inLeft), buildTreeDFS(preorder, memo, inIdx-inLeft+preLeft+1, preRight, inIdx+1)}
}

func buildTreeDFS_(preorder []int, inorder []int, memo map[int]int, preLeft int, preRight int, inLeft int, inRight int) *TreeNode {
	if preLeft > preRight || inLeft > inRight {
		return nil
	}
	//if preRight-preLeft != inRight-inLeft { // err
	//}
	inIdx := memo[preorder[preLeft]]
	curr := &TreeNode{preorder[preLeft], nil, nil}
	curr.Left = buildTreeDFS_(preorder, inorder, memo, preLeft+1, inIdx-inLeft+preLeft, inLeft, inIdx-1) // 可以看出，这里传 inorder inRight 都没用
	curr.Right = buildTreeDFS_(preorder, inorder, memo, inIdx-inLeft+preLeft+1, preRight, inIdx+1, inRight)
	fmt.Println(curr.Val)
	return curr
}

//leetcode submit region end(Prohibit modification and deletion)
