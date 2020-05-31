//给定一个 N 叉树，返回其节点值的层序遍历。 (即从左到右，逐层遍历)。 
//
// 例如，给定一个 3叉树 : 
//
// 
//
// 
//
// 
//
// 返回其层序遍历: 
//
// [
//     [1],
//     [3,2,4],
//     [5,6]
//]
// 
//
// 
//
// 说明: 
//
// 
// 树的深度不会超过 1000。 
// 树的节点总数不会超过 5000。 
// Related Topics 树 广度优先搜索
package main

import "fmt"

func main() {
	//root := &Node{1, []*Node{{3, nil}, nil, {4, nil}}}
	//fmt.Println(root.Children)	// 包含 nil：[0xc0000044c0 <nil> 0xc0000044e0]
	root := &Node{1, []*Node{{3, []*Node{{5, nil}, {6, nil}}}, {2, nil}, {4, nil}}}
	order := levelOrder(root)
	fmt.Println(order) // [[1], [3, 2, 4], [5, 6]]
}

type Node struct {
	Val      int
	Children []*Node
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	// Stack
	order := make([][]int, 0)
	if root == nil {
		return order
	}
	stack := make([]*Node, 1)
	stack[0] = root
	for len(stack) != 0 {
		ord := make([]int, 0)
		n := len(stack)
		for i := 0; i < n; i++ {
			//if stack[i] != nil {
				ord = append(ord, stack[i].Val)
				stack = append(stack, stack[i].Children...)
			//}
		}
		//for _, node := range stack {
		//	ord = append(ord, node.Val)
		//	stack = append(stack, node.Children...)
		//}
		order = append(order, ord)
		stack = stack[n:]
	}
	return order

	// 递归
	//order := make([][]int, 0)
	//levelOrderRecursion(root, 0, &order)
	//return order
}

func levelOrderRecursion(root *Node, level int, order *[][]int) {
	if root == nil {
		return
	}
	if level == len(*order) {
		*order = append(*order, []int{})
	}
	(*order)[level] = append((*order)[level], root.Val)
	for _, node := range root.Children {
		levelOrderRecursion(node, level+1, order)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
