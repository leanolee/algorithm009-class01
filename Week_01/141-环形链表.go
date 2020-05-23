//给定一个链表，判断链表中是否有环。 
//
// 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。 
//
// 
//
// 示例 1： 
//
// 输入：head = [3,2,0,-4], pos = 1
//输出：true
//解释：链表中有一个环，其尾部连接到第二个节点。
// 
//
// 
//
// 示例 2： 
//
// 输入：head = [1,2], pos = 0
//输出：true
//解释：链表中有一个环，其尾部连接到第一个节点。
// 
//
// 
//
// 示例 3： 
//
// 输入：head = [1], pos = -1
//输出：false
//解释：链表中没有环。
// 
//
// 
//
// 
//
// 进阶： 
//
// 你能用 O(1)（即，常量）内存解决此问题吗？ 
// Related Topics 链表 双指针
package main

import "fmt"

func main() {
	head := &ListNode{3, nil}
	node1 := &ListNode{2, nil}
	node2 := &ListNode{0, nil}
	node3 := &ListNode{-4, nil}
	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node1
	cycle := hasCycle(head)
	fmt.Println(cycle)
}

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
/*
暴力法：O(n^2),O(n)
	思路：每过一个节点，将此节点放入数组中，并遍历数组，找到相同的节点，则返回true
hash：O(n),O(n)
	思路：将数组换为 hash
快慢指针：O(n),O(1)
	更多请见 leetcode-142
	1.slow、fast 指针，每次分别走 1、2 步
	2.当 slow、fast 的节点相同时，返回true
*/
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	// 快慢指针
	if head == nil {
		return false
	}
	fast := head.Next
	//for head != nil && fast != nil && fast.Next != nil {
	for fast != nil && fast.Next != nil {
		if head == fast {
			return true
		}
		head = head.Next
		fast = fast.Next.Next
	}
	return false

	// 暴力
	//memo := make(map[*ListNode]int)
	//for ; head != nil; head = head.Next {
	//	if _, ok := memo[head]; ok {
	//		return true
	//	}
	//	memo[head] = 1
	//}
	//return false

	// 更暴力
	//memo := make([]*ListNode, 0)
	//for ; head != nil; head = head.Next {
	//	memo = append(memo, head)
	//	for _, node := range memo {
	//		if head == node {
	//			return true
	//		}
	//	}
	//}
	//return false
}

//leetcode submit region end(Prohibit modification and deletion)
