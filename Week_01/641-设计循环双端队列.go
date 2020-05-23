//设计实现双端队列。 
//你的实现需要支持以下操作： 
//
// 
// MyCircularDeque(k)：构造函数,双端队列的大小为k。 
// insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true。 
// insertLast()：将一个元素添加到双端队列尾部。如果操作成功返回 true。 
// deleteFront()：从双端队列头部删除一个元素。 如果操作成功返回 true。 
// deleteLast()：从双端队列尾部删除一个元素。如果操作成功返回 true。 
// getFront()：从双端队列头部获得一个元素。如果双端队列为空，返回 -1。 
// getRear()：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1。 
// isEmpty()：检查双端队列是否为空。 
// isFull()：检查双端队列是否满了。 
// 
//
// 示例： 
//
// MyCircularDeque circularDeque = new MycircularDeque(3); // 设置容量大小为3
//circularDeque.insertLast(1);			        // 返回 true
//circularDeque.insertLast(2);			        // 返回 true
//circularDeque.insertFront(3);			        // 返回 true
//circularDeque.insertFront(4);			        // 已经满了，返回 false
//circularDeque.getRear();  				// 返回 2
//circularDeque.isFull();				        // 返回 true
//circularDeque.deleteLast();			        // 返回 true
//circularDeque.insertFront(4);			        // 返回 true
//circularDeque.getFront();				// 返回 4
//  
//
// 
//
// 提示： 
//
// 
// 所有值的范围为 [1, 1000] 
// 操作次数的范围为 [1, 1000] 
// 请不要使用内置的双端队列库。 
// 
// Related Topics 设计 队列
package main

import (
	"fmt"
)

func main() {
	myCircularDeque := Constructor(5)
	fmt.Println(myCircularDeque)
}
/*
使用链表：
	题中所涉及方法都是：O(1)
使用数组：仔细品味
	题中所涉及方法都是：O(1)
	1.需要5个成员
		type MyCircularDeque struct {
			deque []int
			len   int
			cap   int
			head  int
			tail  int
		}
	2.固定了cap后，通过索引记录当前数组的：head、tail
		则在原数组空间上，即可完成需求
*/
//leetcode submit region begin(Prohibit modification and deletion)
// 使用数组
type MyCircularDeque struct {
	deque []int
	len   int
	cap   int
	head  int
	tail  int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{make([]int, k, k), 0, k, 0, k - 1}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.len++
	this.head--
	if this.head == -1 {
		this.head = this.cap - 1
	}
	this.deque[this.head] = value
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.len++
	this.tail++
	if this.tail == this.cap {
		this.tail = 0
	}
	this.deque[this.tail] = value
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.len--
	this.head++
	if this.head == this.cap {
		this.head = 0
	}
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.len--
	this.tail--
	if this.tail == -1 {
		this.tail = this.cap - 1
	}
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.deque[this.head]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.deque[this.tail]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.len == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.len == this.cap
}

// 使用链表
//type MyCircularDeque struct {
//	head   *Node
//	tail   *Node
//	length int
//	cap    int
//}
//type Node struct {
//	pre   *Node
//	next  *Node
//	value int
//}
//
///** Initialize your data structure here. Set the size of the deque to be k. */
//func Constructor(k int) MyCircularDeque {
//	return MyCircularDeque{nil, nil, 0, k}
//}
//
///** Adds an item at the front of Deque. Return true if the operation is successful. */
//func (this *MyCircularDeque) InsertFront(value int) bool {
//	if this.length == this.cap {
//		return false
//	}
//	node := &Node{nil, nil, value}
//	if this.length == 0 {
//		this.head = node
//		this.tail = node
//	} else {
//		this.head.pre = node
//		node.next = this.head
//		this.head = node
//	}
//	this.length++
//	return true
//}
//
///** Adds an item at the rear of Deque. Return true if the operation is successful. */
//func (this *MyCircularDeque) InsertLast(value int) bool {
//	if this.length == this.cap {
//		return false
//	}
//	node := &Node{nil, nil, value}
//	if this.length == 0 {
//		this.head = node
//		this.tail = node
//	} else {
//		this.tail.next = node
//		node.pre = this.tail
//		this.tail = node
//	}
//	this.length++
//	return true
//}
//
///** Deletes an item from the front of Deque. Return true if the operation is successful. */
//func (this *MyCircularDeque) DeleteFront() bool {
//	if this.length == 0 {
//		return false
//	}
//	node := this.head.next
//	if node == nil {
//		this.head = nil
//		this.tail = nil
//	} else {
//		this.head = node
//		node.pre = nil
//	}
//	this.length--
//	return true
//}
//
///** Deletes an item from the rear of Deque. Return true if the operation is successful. */
//func (this *MyCircularDeque) DeleteLast() bool {
//	if this.length == 0 {
//		return false
//	}
//	node := this.tail.pre
//	if node == nil {
//		this.head = nil
//		this.tail = nil
//	} else {
//		this.tail = node
//		node.next = nil
//	}
//	this.length--
//	return true
//}
//
///** Get the front item from the deque. */
//func (this *MyCircularDeque) GetFront() int {
//	if this.length == 0 {
//		return -1
//	}
//	return this.head.value
//}
//
///** Get the last item from the deque. */
//func (this *MyCircularDeque) GetRear() int {
//	if this.length == 0 {
//		return -1
//	}
//	return this.tail.value
//}
//
///** Checks whether the circular deque is empty or not. */
//func (this *MyCircularDeque) IsEmpty() bool {
//	return this.length == 0
//}
//
///** Checks whether the circular deque is full or not. */
//func (this *MyCircularDeque) IsFull() bool {
//	return this.length == this.cap
//}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
//leetcode submit region end(Prohibit modification and deletion)
