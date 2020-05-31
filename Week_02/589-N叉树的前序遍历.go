//给定一个 N 叉树，返回其节点值的前序遍历。 
//
// 例如，给定一个 3叉树 : 
//
// 
//
// 
//
// 
//
// 返回其前序遍历: [1,3,5,6,2,4]。 
//
// 
//
// 说明: 递归法很简单，你可以使用迭代法完成此题吗? Related Topics 树
package main

import "fmt"

func main() {
	root := &Node{1, []*Node{{3, []*Node{{5, nil}, {6, nil}}}, {2, nil}, {4, nil}}}
	ints := preorder(root)
	fmt.Println(ints) // [1,3,5,6,2,4]
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

func preorder(root *Node) []int {
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
		for i := len(curr.Children) - 1; i >= 0; i-- {
			stack = append(stack, curr.Children[i])
		}
	}
	return ints

	// 递归
	//ints := make([]int, 0)
	//preorderRecursion(root, &ints)
	//return ints
}

func preorderRecursion(root *Node, ints *[]int) {
	if root == nil {
		return
	}
	*ints = append(*ints, root.Val)
	for _, node := range root.Children {
		preorderRecursion(node, ints)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
