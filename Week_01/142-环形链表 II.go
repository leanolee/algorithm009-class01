//给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。 
//
// 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。 
//
// 说明：不允许修改给定的链表。 
//
// 
//
// 示例 1： 
//
// 输入：head = [3,2,0,-4], pos = 1
//输出：tail connects to node index 1
//解释：链表中有一个环，其尾部连接到第二个节点。
// 
//
// 
//
// 示例 2： 
//
// 输入：head = [1,2], pos = 0
//输出：tail connects to node index 0
//解释：链表中有一个环，其尾部连接到第一个节点。
// 
//
// 
//
// 示例 3： 
//
// 输入：head = [1], pos = -1
//输出：no cycle
//解释：链表中没有环。
// 
//
// 
//
// 
//
// 进阶： 
//你是否可以不用额外空间解决此题？ 
// Related Topics 链表 双指针
package main

import "fmt"

func main() {
	head := &ListNode{3, nil}
	node1 := &ListNode{2, nil}
	head.Next = node1

	node2 := &ListNode{0, nil}
	node3 := &ListNode{-4, nil}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node1

	//node1.Next = head
	cycle := detectCycle(head)
	fmt.Println(cycle)
}

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
/*
暴力hash：O(n),O(n)
	同 leetcode-141
快慢指针：O(n),O(1)
	规律：
		快指针每次两步，慢指针每次一步。快指针每跑满2圈，慢指针跑满一圈，相遇在上次相遇的同一个点
		另外，第2次相遇的步数-第一次相遇的步数，等于环形链表中的节点数
	题解：
		假设链表总共 a + b 个节点，a表示环形之外的节点数，b表示环形内的节点数。有slow、fast两个指针
		第一次相遇：slow、fast走过的节点数关系：
			fast = 2*slow = 2s，fast - slow = nb（n个b），则 slow=s=nb，fast=2nb
		第二次相遇：fast变成和slow一样，每次一步，则有：
			fast = a 时，slow = a + nb，两者刚好在【入环的第一个节点】相遇
*/
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	// 快慢指针：优化
	//slow, fast := head, head
	//for fast != nil && fast.Next != nil {
	//	slow = slow.Next
	//	fast = fast.Next.Next
	//	if slow == fast {
	//		for slow != head {
	//			slow = slow.Next
	//			head = head.Next
	//		}
	//		return head
	//	}
	//}
	//return nil

	// 快慢指针
	slow, fast := head, head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	fast = head
	for {
		if slow == fast {
			return fast
		}
		slow = slow.Next
		fast = fast.Next
	}

	// 快慢指针：失败
	//if head == nil {
	//	return nil
	//}
	//fast := head
	//first, second := 0, 0
	//for fast != nil && fast.Next != nil {
	//	head = head.Next
	//	fast = fast.Next.Next
	//	second++
	//	if fast != head {
	//		continue
	//	}
	//	// 这样行不通，无法计算a、b值
	//	if first > 0 {
	//		for i := 0; i < second-first+1; i++ {
	//			head = head.Next
	//		}
	//		return head
	//	}
	//	first = second
	//}
	//return nil

	// 暴力
	//memo := make(map[*ListNode]int)
	//for ; head != nil; head = head.Next {
	//	if _, ok := memo[head]; ok {
	//		return head
	//	}
	//	memo[head] = 1
	//}
	//return nil
}

//leetcode submit region end(Prohibit modification and deletion)
