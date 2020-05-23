//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。 
//
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。 
//
// 
//
// 示例: 
//
// 给定 1->2->3->4, 你应该返回 2->1->4->3.
// 
// Related Topics 链表
package main

import (
	"fmt"
)

func main() {
	head := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	pairs := swapPairs(&head)
	for next := pairs; next != nil; next = next.Next {
		fmt.Println(next.Val)
	}
}

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
暴力法：O(n),O(n)
	1.将链表转为数组
	2.一次循环：判断索引
		为奇数，通过中间变量 node 记录该节点，按需求拼接其他节点
		为偶数，直接添加该节点
	3.如果遍历完，node != nil，接入该节点，并将其next置为nil
递归：O(n/2)=O(n),O(1)
	1.终结条件
		head==nil || head.next==nil
	2.逻辑处理
		按需交换节点，golang可一次交换多个值
	3.下一层
		head.next.next进入下一层
	4.清理当前层
迭代：O(n),O(1)
	一次循环：利用中间变量来交换节点
迭代优化：O(n),O(1)
	一次循环：一次性交换多个节点
 */
func swapPairs(head *ListNode) *ListNode {
	// leetcode精妙办法
	//temp := &ListNode{0, head}
	//newHead := temp
	////fmt.Println(temp, temp.Next, temp.Next.Next, temp.Next.Next.Next)
	//for temp.Next != nil && temp.Next.Next != nil {
	//	temp.Next, temp.Next.Next.Next, temp.Next.Next, temp = temp.Next.Next, temp.Next, temp.Next.Next.Next, temp.Next
	//	//fmt.Println(newHead, newHead.Next, newHead.Next.Next, newHead.Next.Next.Next)
	//}
	//return newHead.Next

	// 递归
	//if head == nil || head.Next == nil {
	//	return head
	//}
	//head, head.Next.Next, head.Next = head.Next, head, swapPairs(head.Next.Next)
	//return head

	// 笨办法2
	//newHead := &ListNode{0, head}
	//temp := newHead
	//var node *ListNode
	//for temp.Next != nil && temp.Next.Next != nil {
	//	node = temp.Next
	//	temp.Next = temp.Next.Next
	//	node.Next = temp.Next.Next
	//	temp.Next.Next = node
	//	temp = node
	//}
	//return newHead.Next

	// 笨办法
	//if head == nil || head.Next == nil {
	//	return head
	//}
	//nodes := make([]*ListNode, 0)
	//for next := head; next != nil; next = next.Next {
	//	nodes = append(nodes, next)
	//}
	//var node, nodeEnd, newNode *ListNode
	//for i := 0; i < len(nodes); i++ {
	//	if i&1 == 1 {
	//		if newNode == nil {
	//			newNode = nodes[i]
	//			nodeEnd = newNode
	//		} else {
	//			nodeEnd.Next = nodes[i]
	//			nodeEnd = nodes[i]
	//		}
	//		if node != nil && newNode != nil {
	//			nodeEnd.Next = node
	//			nodeEnd = node
	//			node = nil
	//		}
	//	} else {
	//		node = nodes[i]
	//	}
	//}
	//if node != nil {
	//	nodeEnd.Next = node
	//	nodeEnd = node
	//}
	//nodeEnd.Next = nil
	//return newNode
}

//leetcode submit region end(Prohibit modification and deletion)
