//输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。 
//
// 
//
// 示例 1： 
//
// 输入：head = [1,3,2]
//输出：[2,3,1] 
//
// 
//
// 限制： 
//
// 0 <= 链表长度 <= 10000 
//
package main

func main() {

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
func reversePrint(head *ListNode) []int {
	list := make([]int, 0)
	reversePrintDFS(head, &list)
	return list
}

func reversePrintDFS(head *ListNode, list *[]int) {
	// Stack
	// 递归
	if head == nil {
		return
	}
	reversePrintDFS(head.Next, list)
	*list = append(*list, head.Val)
}

//leetcode submit region end(Prohibit modification and deletion)
