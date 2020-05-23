//给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。 
//
// k 是一个正整数，它的值小于或等于链表的长度。 
//
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。 
//
// 
//
// 示例： 
//
// 给你这个链表：1->2->3->4->5 
//
// 当 k = 2 时，应当返回: 2->1->4->3->5 
//
// 当 k = 3 时，应当返回: 3->2->1->4->5 
//
// 
//
// 说明： 
//
// 
// 你的算法只能使用常数的额外空间。 
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。 
// 
// Related Topics 链表
package main

import (
	"fmt"
)

func main() {
	//head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	head := &ListNode{1, &ListNode{2, nil}}
	k := 2
	group := reverseKGroup(head, k)
	for ; group != nil; group = group.Next {
		fmt.Println(group.Val)
	}
	//reverse := reverseNode(head)
	//for ; reverse != nil; reverse = reverse.Next {
	//	fmt.Println(reverse.Val)
	//}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
/*
递归：O(n),O(1)
	1.终止条件
		head==nil
	2.逻辑处理
		一次循环：取 k 个节点，进行reverse
	3.下一层
		将剩余节点，递归到下一层
		如果上次翻转的数量 count < k，则再将这些节点进行一次递归，但 k=count
	4.清理当前层
Stack：O(n),O(k)
	1.创建大小为 k 的 stack
	2.一次循环：节点压入栈中直到 k 个，再 for 循环依次从 stack 中弹出
	3.如果stack中还有值，再翻转回原顺序
Stack优化：O(n),O(n)
	1.一次循环：创建 memo，用于存储所有节点的 stack
	2.计算：memo.len % k，用于截取掉不用翻转节点
	3.二次循环：memo中，每 k 个元素进行翻转，countK:0~k-1
		i+k-1-countK<<1：当前节点索引的算法
	4.将剩余的节点，按原顺序加入
迭代：O(n),O(1)
	1.需要使用两个变量 start、newStart 分别记录：每次翻转遍历的起始节点、当前翻转结束节点的下一个节点
	2.一层循环：遍历原链表（虽然是2层循环，但是两两是独立，所以仍然是O(n)复杂度
		每 k 个节点，进行翻转
	3.二层循环：数 k(k-1) 个节点，如果不够 k 个，则直接原顺序接入
*/
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseNode(reverse *ListNode) *ListNode {
	var nodesK *ListNode
	for ; reverse != nil; {
		reverse.Next, nodesK, reverse = nodesK, reverse, reverse.Next
	}
	return nodesK
}
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 迭代
	newHead := new(ListNode)
	pre := newHead
	var start, newStart *ListNode // start：每次翻转遍历的起始节点；newStart：结束节点的下一个节点
	for head != nil {
		start = head
		for i := 1; i < k; i++ {
			head = head.Next
			if head == nil {
				pre.Next = start
				return newHead.Next
			}
		}
		newStart, head.Next = head.Next, nil
		pre.Next = reverseNode(start)
		pre, head = start, newStart
	}
	return newHead.Next

	// 栈：优化
	//memo := make([]*ListNode, 0)
	//newHead := new(ListNode)
	//temp := newHead
	//for head != nil {
	//	memo = append(memo, head)
	//	head = head.Next
	//}
	//count := len(memo)
	//positive := count % k
	//countK := 0
	//for i := 0; i < count-positive; i++ {
	//	if countK == k {
	//		countK = 0
	//	}
	//	// 不能合并，因为初始化temp时，是个单独的节点
	//	temp.Next = memo[i+k-1-countK<<1]
	//	temp = temp.Next
	//	if countK == k-1 {
	//		temp.Next = nil
	//	}
	//	countK++
	//}
	//for i := positive; i > 0; i-- {
	//	temp.Next = memo[count-i]
	//	temp = temp.Next
	//}
	//return newHead.Next

	//栈
	//memo := make([]*ListNode, k)
	//count := 0
	//var start, newHead *ListNode
	//isFirst := true
	//for head != nil {
	//	memo[count], head = head, head.Next
	//	count++
	//	if count == k {
	//		count = 0
	//		for i := k - 1; i > 0; i-- {
	//			memo[i].Next = memo[i-1]
	//		}
	//		if isFirst {
	//			isFirst = false
	//			newHead = memo[k-1]
	//		} else {
	//			start.Next = memo[k-1]
	//		}
	//		start = memo[0]
	//		start.Next = nil
	//	}
	//}
	//for i := 0; i < count; i++ {
	//	start.Next, start = memo[i], memo[i]
	//}
	//return newHead

	// 递归 + 迭代
	//return reverseKGroupRecursion(nil, 0, head, head, k)
}

func reverseKGroupRecursion(newHead *ListNode, count int, start *ListNode, head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	for head != nil {
		head, head.Next, newHead = head.Next, newHead, head
		count++
		if count == k {
			start.Next = reverseKGroupRecursion(nil, 0, head, head, k)
			return newHead
		}
	}
	return reverseKGroupRecursion(nil, 0, newHead, newHead, count)
}

//leetcode submit region end(Prohibit modification and deletion)
