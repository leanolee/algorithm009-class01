//将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
//
// 示例： 
//
// 输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
// 
// Related Topics 链表
package main

import "fmt"

func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	lists := mergeTwoLists(l1, l2)
	for ; lists != nil; lists = lists.Next {
		fmt.Println(lists.Val)
	}
}

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
/*
递归：O(m+n),O(1)
	1.terminator
		如果其中一个链表==nil，将另外一个链表接到返回值中，返回
	2.process result
		比较两个链表当前的 value 大小
	3.process current logic
		将 小 的那个链表的当前节点接入总链中，且将递归的下一个总链置为该 小 链表，l.next 进下一层
	4.restore current status
迭代：O(m+n),O(1)
	1.循环 l1,l2，直到都为nil
	2.如果其中之一为nil，拼接另外一个
	3.将较小的拼接到结果中
*/
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 暴力
	list := new(ListNode)
	temp := list
	one, two := 0, 0
	for l1 != nil || l2 != nil {
		if l1 == nil {
			temp.Next = l2
			break
		}
		if l2 == nil {
			temp.Next = l1
			break
		}
		one, two = l1.Val, l2.Val
		if one < two {
			temp.Next = l1
			temp = l1
			l1 = l1.Next
		} else {
			temp.Next = l2
			temp = l2
			l2 = l2.Next
		}
	}
	return list.Next

	// 递归
	//list := new(ListNode)
	//mergeTwoListsRecursion(l1, l2, list)
	//return list.Next
}

func mergeTwoListsRecursion(l1 *ListNode, l2 *ListNode, list *ListNode) {
	if l1 == nil {
		list.Next = l2
		return
	}
	if l2 == nil {
		list.Next = l1
		return
	}
	one, two := l1.Val, l2.Val
	if one < two {
		list.Next = l1
		mergeTwoListsRecursion(l1.Next, l2, l1)
	} else {
		list.Next = l2
		mergeTwoListsRecursion(l1, l2.Next, l2)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
