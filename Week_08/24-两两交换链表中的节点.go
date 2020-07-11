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

import "fmt"

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	pairs := swapPairs(head)
	for curr := pairs; curr != nil; curr = curr.Next {
		fmt.Print(curr.Val, " ")
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	// 迭代
	root := &ListNode{0, head}
	pre := root
	for head != nil && head.Next != nil {
		//temp.Next, temp.Next.Next.Next, temp.Next.Next, temp = temp.Next.Next, temp.Next, temp.Next.Next.Next, temp.Next
		pre, pre.Next, head, head.Next, head.Next.Next = head, head.Next, head.Next.Next, head.Next.Next, head
	}
	return root.Next

	// 递归
	//if head == nil || head.Next == nil {
	//	return head
	//}
	//head.Next, head.Next.Next, head = swapPairs(head.Next.Next), head, head.Next
	//return head
}

//leetcode submit region end(Prohibit modification and deletion)
