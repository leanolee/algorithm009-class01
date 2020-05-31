//给定一个 N 叉树，找到其最大深度。 
//
// 最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。 
//
// 例如，给定一个 3叉树 : 
//
// 
//
// 
//
// 
//
// 我们应返回其最大深度，3。 
//
// 说明: 
//
// 
// 树的深度不会超过 1000。 
// 树的节点总不会超过 5000。 
// Related Topics 树 深度优先搜索 广度优先搜索
package main

import (
	"fmt"
)

func main() {
	root := &Node{1, []*Node{{3, []*Node{{5, nil}, {6, nil}}}, {2, nil}, {4, nil}}}
	depth := maxDepth(root)
	fmt.Println(depth)
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

func maxDepth(root *Node) int {
	// BFS
	level := 0
	if root == nil {
		return level
	}
	stack := []*Node{root}
	for len(stack) != 0 {
		level++
		n := len(stack)
		for _, node := range stack {
			//for _, no := range node.Children {
			//	if no != nil {
			//		stack = append(stack, no)
			//	}
			//}
			stack = append(stack, node.Children...)
		}
		stack = stack[n:]
	}
	return level

	// 递归：DFS
	//if root == nil {
	//	return 0
	//}
	//max := 0
	//for _, node := range root.Children {
	//	max = MyTools.Max(maxDepth(node), max)
	//}
	//return max + 1
}

//leetcode submit region end(Prohibit modification and deletion)
