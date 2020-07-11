//运用你所掌握的数据结构，设计和实现一个 LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。 
//
// 获取数据 get(key) - 如果关键字 (key) 存在于缓存中，则获取关键字的值（总是正数），否则返回 -1。 
//写入数据 put(key, value) - 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字/值」。当缓存容量达到上限时，它应该在
//写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。 
//
// 
//
// 进阶: 
//
// 你是否可以在 O(1) 时间复杂度内完成这两种操作？ 
//
// 
//
// 示例: 
//
// LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
//
//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // 返回  1
//cache.put(3, 3);    // 该操作会使得关键字 2 作废
//cache.get(2);       // 返回 -1 (未找到)
//cache.put(4, 4);    // 该操作会使得关键字 1 作废
//cache.get(1);       // 返回 -1 (未找到)
//cache.get(3);       // 返回  3
//cache.get(4);       // 返回  4
// 
// Related Topics 设计
package main

import (
	"fmt"
	"math"
)

func main() {
	//lru := Constructor(1)
	//lru.Put(2, 1)
	//get := lru.Get(2)
	//fmt.Println(get)
	//lru.Put(3, 2)
	//get = lru.Get(2)
	//fmt.Println(get)
	//get = lru.Get(3)
	//fmt.Println(get)

	lru := Constructor(2)
	lru.Put(2, 1)
	lru.Put(3, 2)
	get := lru.Get(3)
	fmt.Println(get)
	get = lru.Get(2)
	fmt.Println(get)
	lru.Put(4, 3)
	get = lru.Get(2)
	fmt.Println(get)
	//for k, v := range lru.cache {
	//	fmt.Println(k, v.val, v.pre.key)
	//}
	get = lru.Get(3)
	fmt.Println(get)
	get = lru.Get(4)
	fmt.Println(get)
}

/*
hash + 双向链表
	hash：get O(1)
	双向链表:put O(1)
*/
//leetcode submit region begin(Prohibit modification and deletion)
// lc
type LRUCache struct {
	size  int
	cap   int
	cache map[int]*DeList
	head  *DeList
	tail  *DeList
}

func Constructor(capacity int) LRUCache {
	head, tail := &DeList{math.MinInt64, math.MinInt64, nil, nil}, &DeList{math.MaxInt64, math.MaxInt64, nil, nil}
	head.next, tail.pre = tail, head
	return LRUCache{0, capacity, make(map[int]*DeList), head, tail}
}
func (this *LRUCache) Get(key int) int { // 重点：map中添加的是节点的指针，所以修改了node不用再重新放入map
	deList, ok := this.cache[key]
	if !ok { // 1.不存在
		return -1
	}
	if deList.pre != this.head { // 3.不在首位
		this.removeNode(deList) // 3.1.移除
		this.addHead(deList)    // 3.2.添加到首
	}
	return deList.val // 2.存在
}
func (this *LRUCache) Put(key int, value int) { // 重点：map中添加的是节点的指针，所以修改了node不用再重新放入map
	deList, ok := this.cache[key]
	if ok { // 1.已存在
		deList.val = value
		if deList.pre == this.head { // 1.1.已经在首位
			return
		}
		this.removeNode(deList) // 1.2.移除
	} else {                       // 2.不存在
		if this.size == this.cap { // 2.1.超出 cap，移除尾节点
			delete(this.cache, this.removeTail().key)
			this.size--
		}
		deList = &DeList{key, value, nil, nil} // 2.2.构建新的节点，并存入map
		this.cache[key] = deList
		this.size++
	}
	this.addHead(deList) // 3.添加到首
}
func (this *LRUCache) addHead(node *DeList) { // 添加到首节点
	node.pre, node.next, this.head.next, this.head.next.pre = this.head, this.head.next, node, node
}
func (this *LRUCache) removeNode(node *DeList) { // 移除一个节点
	node.pre.next, node.next.pre = node.next, node.pre
}
func (this *LRUCache) removeTail() *DeList { // 移除尾节点
	node := this.tail.pre
	this.removeNode(node)
	return node
}

type DeList struct {
	key, val  int
	pre, next *DeList
}

// 个人写法：重点：map中添加的是节点的指针，所以修改了node不用再重新放入map
//type LRUCache struct {
//	size  int
//	cap   int
//	cache map[int]*DeList
//	head  *DeList
//	tail  *DeList
//}
//
//func Constructor(capacity int) LRUCache {
//	head, tail := &DeList{math.MinInt64, math.MinInt64, nil, nil}, &DeList{math.MaxInt64, math.MaxInt64, nil, nil}
//	head.next, tail.pre = tail, head
//	return LRUCache{0, capacity, make(map[int]*DeList), head, tail}
//}
//
//func (this *LRUCache) Get(key int) int {
//	deList, ok := this.cache[key]
//	if !ok { // 1.不存在
//		return -1
//	}
//	if deList.pre != this.head { // 3.不在首位
//		deList.pre.next, deList.next.pre = deList.next, deList.pre // 3.1.删除
//		this.cache[deList.pre.key], this.cache[deList.next.key] = deList.pre, deList.next
//		this.head.next, this.head.next.pre, deList.pre, deList.next = deList, deList, this.head, this.head.next // 3.2.移到首（变动4个）
//		this.cache[key] = deList
//	}
//	return deList.val // 2.存在
//}
//
//func (this *LRUCache) Put(key int, value int) {
//	deList, ok := this.cache[key]
//	if ok { // 1.已存在
//		deList.val = value
//		if deList.pre == this.head { // 1.1.在首位
//			this.cache[key] = deList
//			return
//		}
//		deList.pre.next, deList.next.pre = deList.next, deList.pre // 1.2.删除
//		this.cache[deList.pre.key], this.cache[deList.next.key] = deList.pre, deList.next
//	} else {                       // 2.不存在
//		if this.size == this.cap { // 2.1.移除尾
//			delete(this.cache, this.tail.pre.key) // a.从 cache 中移除
//			currPre := this.tail.pre.pre
//			this.tail.pre, currPre.next = currPre, this.tail // b.移除最后一个
//			this.cache[currPre.key] = currPre
//			this.size--
//		}
//		deList = &DeList{key, value, nil, nil} // 2.2.构建新的节点
//		this.size++
//	}
//	deList.pre, deList.next, this.head.next, this.head.next.pre = this.head, this.head.next, deList, deList // 3.添到首（变动4个）
//	this.cache[key] = deList
//}
//
//type DeList struct {
//	key, val  int
//	pre, next *DeList
//}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
