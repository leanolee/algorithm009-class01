//给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。 
//
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（
//一个节点也可以是它自己的祖先）。” 
//
// 例如，给定如下二叉树: root = [3,5,1,6,2,0,8,null,null,7,4] 
//
// 
//
// 
//
// 示例 1: 
//
// 输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
//输出: 3
//解释: 节点 5 和节点 1 的最近公共祖先是节点 3。
// 
//
// 示例 2: 
//
// 输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
//输出: 5
//解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。
// 
//
// 
//
// 说明: 
//
// 
// 所有节点的值都是唯一的。 
// p、q 为不同节点且均存在于给定的二叉树中。 
// 
// Related Topics 树
package main

import "fmt"

func main() {
	root := &TreeNode{3, nil, nil}
	node1 := &TreeNode{5, nil, nil}
	node2 := &TreeNode{1, nil, nil}
	node3 := &TreeNode{6, nil, nil}
	node4 := &TreeNode{2, nil, nil}
	node5 := &TreeNode{0, nil, nil}
	node6 := &TreeNode{8, nil, nil}
	node7 := &TreeNode{7, nil, nil}
	node8 := &TreeNode{4, nil, nil}
	root.Left, root.Right = node1, node2
	node1.Left, node1.Right = node3, node4
	node4.Left, node4.Right = node7, node8
	node2.Left, node2.Right = node5, node6
	p, q := node1, node2

	//root := &TreeNode{3, nil, nil}
	//node1 := &TreeNode{5, nil, nil}
	//root.Left = node1
	//p, q := root, node1
	ancestor := lowestCommonAncestor(root, p, q)
	fmt.Println(ancestor)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
重点：
dfs：
	思路：两种情况
		1.p、q 其中之一为另外一个的子树节点
		2.p、q 分别为某一共同节点的左、右子树中的节点
	若 dfs 遍历到当前节点 root == q || root == p，返回
		情况1：先经历一系列 左右节点返回值 分别是 rootL/rootR vs nil 的情况
			最终在最近公共父节点处：rootL vs rootR
		情况2：只会经历 左右节点返回值 是 rootL/rootR vs nil 的情况
存储父节点：
*/
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 存储父节点
	//memo := make(map[*TreeNode]*TreeNode)
	//dfsMemo(memo, root)	// 将所有节点映射为 node->parent
	//parents := make(map[*TreeNode]bool)
	//for upNode := p; upNode != nil; upNode = memo[upNode] {
	//	parents[upNode] = true	// 取出 p 节点的所有父节点
	//}
	//for upNode := q; upNode != nil; upNode = memo[upNode] {
	//	if parents[upNode] {	// 如果 q 的父节点存在于 p 的父节点集合中，返回
	//		return upNode
	//	}
	//}
	//return nil

	// dfs：递归的典范题目
	//if root == nil || root == p || root == q {
	//	return root
	//}
	//left := lowestCommonAncestor(root.Left, p, q)
	//right := lowestCommonAncestor(root.Right, p, q)
	//if left == nil {
	//	return right
	//}
	//if right == nil {
	//	return left
	//}
	//return root

	// BFS：2 超时。问题，索引太大，没有数字可存
	memo := map[int64]*TreeNode{0: root}
	return recursionLowestCommonAncestor([]int64{}, memo, root, p, q)

	// BFS：1 超时。因为数据太大
	//memo := []*TreeNode{root}
	//return lowestCommonAncestorRecursion([]int{}, memo, root, p, q)
}

func recursionLowestCommonAncestor(idx []int64, memo map[int64]*TreeNode, root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	for k, v := range memo {
		if v == p || v == q {
			if !(len(idx) == 1 && idx[0] == k) {
				idx = append(idx, k)
			}
			if len(idx) == 2 {
				return memo[getIdxLowestCommonAncestor(idx)]
			}
		}
		if v.Left != nil {
			memo[k<<1+1] = v.Left
		}
		if v.Right != nil {
			memo[k<<1+2] = v.Right
		}
	}
	return recursionLowestCommonAncestor(idx, memo, root, p, q)
}

func dfsMemo(memo map[*TreeNode]*TreeNode, root *TreeNode) {
	if root.Left != nil {
		memo[root.Left] = root
		dfsMemo(memo, root.Left)
	}
	if root.Right != nil {
		memo[root.Right] = root
		dfsMemo(memo, root.Right)
	}
}

func lowestCommonAncestorRecursion(idx []int, memo []*TreeNode, root, p, q *TreeNode) *TreeNode {
	n := len(memo)
	for i := n >> 1; i < n; i++ {
		node := memo[i]
		if node == nil {
			memo = append(memo, nil, nil)
		} else {
			if node == p || node == q {
				idx = append(idx, i)
				if len(idx) == 2 {
					return memo[getLowestCommonAncestor(idx)]
				}
			}
			memo = append(memo, node.Left, node.Right)
		}
	}
	return lowestCommonAncestorRecursion(idx, memo, root, p, q)
}

func getLowestCommonAncestor(idx []int) int {
	fIdx, sIdx := idx[0], idx[1]
	for fIdx != sIdx {
		if fIdx > sIdx {
			fIdx = (fIdx - 1) >> 1
		} else {
			sIdx = (sIdx - 1) >> 1
		}
	}
	return fIdx
}
func getIdxLowestCommonAncestor(idx []int64) int64 {
	fIdx, sIdx := idx[0], idx[1]
	for fIdx != sIdx {
		if fIdx > sIdx {
			fIdx = (fIdx - 1) >> 1
		} else {
			sIdx = (sIdx - 1) >> 1
		}
	}
	idx = idx[2:]
	return fIdx
}

//leetcode submit region end(Prohibit modification and deletion)
