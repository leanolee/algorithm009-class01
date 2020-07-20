//给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。 
//
// 示例 1: 
//
// 输入: 1->1->2
//输出: 1->2
// 
//
// 示例 2: 
//
// 输入: 1->1->2->3->3
//输出: 1->2->3 
// Related Topics 链表
package main

import (
	"fmt"
)

func main() {
	head := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}
	duplicates := deleteDuplicates(head)
	for curr := duplicates; curr != nil; curr = curr.Next {
		fmt.Println(curr.Val)
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
func deleteDuplicates(head *ListNode) *ListNode {
	pre, curr := head, head
	for curr != nil {
		for curr != nil && curr.Val == pre.Val {
			curr = curr.Next
		}
		pre, pre.Next = curr, curr
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)
