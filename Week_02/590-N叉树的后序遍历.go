//给定一个 N 叉树，返回其节点值的后序遍历。 
//
// 例如，给定一个 3叉树 : 
//
// 
//
// 
//
// 
//
// 返回其后序遍历: [5,6,3,2,4,1]. 
//
// 
//
// 说明: 递归法很简单，你可以使用迭代法完成此题吗? Related Topics 树
package main

import (
	"./MyTools"
	"fmt"
)

func main() {
	root := &Node{1, []*Node{{3, []*Node{{5, nil}, {6, nil}}}, {2, nil}, {4, nil}}}
	ints := postorder(root)
	fmt.Println(ints) // [5 6 3 2 4 1]
}

//type Node struct {
//	Val      int
//	Children []*Node
//}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
	// Stack
	ints := make([]int, 0)
	if root == nil {
		return ints
	}
	stack := make([]*Node, 1)
	stack[0] = root
	var curr *Node
	for len(stack) != 0 {
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ints = append(ints, curr.Val)
		for i := 0; i < len(curr.Children); i++ {
			if curr.Children[i] != nil {
				stack = append(stack, curr.Children[i])
			}
		}
	}
	MyTools.Reverse(ints)
	return ints

	// 递归
	//ints := make([]int, 0)
	//postorderRecursion(root, &ints)
	//return ints
}

func postorderRecursion(root *Node, ints *[]int) {
	if root == nil {
		return
	}
	for _, node := range root.Children {
		postorderRecursion(node, ints)
	}
	*ints = append(*ints, root.Val)
}

//leetcode submit region end(Prohibit modification and deletion)
