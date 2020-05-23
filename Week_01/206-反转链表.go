//反转一个单链表。 
//
// 示例: 
//
// 输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL 
//
// 进阶: 
//你可以迭代或递归地反转链表。你能否用两种方法解决这道题？ 
// Related Topics 链表
package main

import "fmt"

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	list := reverseList(head)
	for ; list != nil; list = list.Next {
		fmt.Println(list.Val)
	}
}

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

/*
递归：O(n),O(1)
	主要是节点的交换，需要画图解
栈：O(n),O(n)
	1.一次遍历：将所有节点压入栈中
	2.二次遍历：倒序遍历数组，拼接链表
迭代：O(n),O(1)
	记录 next 节点
	当前的节点，指向上一个
	新建节点的头 newHead 等于当前节点
*/
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	// 迭代
	var newHead *ListNode
	for head != nil {
		head, head.Next, newHead = head.Next, newHead, head
	}
	return newHead

	// Stack
	//if head == nil {
	//	return head
	//}
	//list := make([]*ListNode, 0)
	//for node := head; node != nil; node = node.Next {
	//	list = append(list, node)
	//}
	//for i := len(list) - 2; i >= 0; i-- {
	//	list[i+1].Next = list[i]
	//}
	//list[0].Next = nil
	//return list[len(list)-1]

	// 递归
	//if head == nil || head.Next == nil {
	//	return head
	//}
	//reverse := reverseList(head.Next)
	//head.Next.Next = head
	//head.Next = nil
	//return reverse
}

//leetcode submit region end(Prohibit modification and deletion)
