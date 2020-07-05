//用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的
//功能。(若队列中没有元素，deleteHead 操作返回 -1 ) 
//
// 
//
// 示例 1： 
//
// 输入：
//["CQueue","appendTail","deleteHead","deleteHead"]
//[[],[3],[],[]]
//输出：[null,null,3,-1]
// 
//
// 示例 2： 
//
// 输入：
//["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
//[[],[],[5],[2],[],[]]
//输出：[null,-1,null,null,5,2]
// 
//
// 提示： 
//
// 
// 1 <= values <= 10000 
// 最多会对 appendTail、deleteHead 进行 10000 次调用 
// 
// Related Topics 栈 设计
package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
type CQueue struct {
	add []int // 在顶 append
	del []int // 在顶 delete
}

func Constructor() CQueue {
	return CQueue{make([]int, 0), make([]int, 0)}
}

func (this *CQueue) AppendTail(value int) {
	if len(this.del) > 0 {
		for i := len(this.del) - 1; i >= 0; i-- {
			this.add = append(this.add, this.del[i])
		}
		this.del = this.del[:0]
	}
	this.add = append(this.add, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.add)+len(this.del) == 0 {
		return -1
	}
	if len(this.add) > 0 {
		for i := len(this.add) - 1; i >= 0; i-- {
			this.del = append(this.del, this.add[i])
		}
		this.add = this.add[:0]
	}
	last := len(this.del) - 1
	val := this.del[last]
	this.del = this.del[:last]
	return val
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
//leetcode submit region end(Prohibit modification and deletion)
