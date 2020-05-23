//设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。 
//
// 
// push(x) —— 将元素 x 推入栈中。 
// pop() —— 删除栈顶的元素。 
// top() —— 获取栈顶元素。 
// getMin() —— 检索栈中的最小元素。 
// 
//
// 
//
// 示例: 
//
// 输入：
//["MinStack","push","push","push","getMin","pop","top","getMin"]
//[[],[-2],[0],[-3],[],[],[],[]]
//
//输出：
//[null,null,null,null,-3,null,0,-2]
//
//解释：
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.getMin();   --> 返回 -2.
// 
//
// 
//
// 提示： 
//
// 
// pop、top 和 getMin 操作总是在 非空栈 上调用。 
// 
// Related Topics 栈 设计
package main

import "fmt"

func main() {
	minStack := Constructor()
	fmt.Println(minStack)
}
/*
栈：
	个人笨方法
	思路：
		1.维护两个栈，分别存所有元素，最小元素
		type MinStack struct {
			stack    []int
			minStack []int
		}
		2.Push：
			先放入stack。如果小于 minStack 栈顶元素，则放入 minStack
		3.Getmin：
			取 minStack 栈顶元素。如果没有，按照 Push 的思路，从 stack 中往 minStack 添加元素
Stack：
	思路：不同之处在于
		1.Push：先放入stack。
			如果小于 minStack 栈顶元素，则放入 minStack
			否则，将 minStack 栈顶元素复制一个，再一次放入 minStack
		2.GetMin：
			每次直接取栈顶元素
*/
//leetcode submit region begin(Prohibit modification and deletion)
// leetcode方法，仔细体会 栈 的精妙
type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, x)
		return
	}
	newMin := this.minStack[len(this.minStack)-1]
	if x < newMin {
		newMin = x
	}
	this.minStack = append(this.minStack, newMin)
}

func (this *MinStack) Pop() {
	this.minStack = this.minStack[:len(this.minStack)-1]
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

// 个人的笨办法
//type MinStack struct {
//	//top    int
//	//bottom int
//	//next   *MinStack
//	//pre    *MinStack
//	stack    []int
//	minStack []int
//}
//
///** initialize your data structure here. */
//func Constructor() MinStack {
//	return MinStack{[]int{}, []int{}}
//}
//
//func (this *MinStack) Push(x int) {
//	this.stack = append(this.stack, x)
//	if len(this.minStack) == 0 || x <= this.minStack[len(this.minStack)-1] {
//		this.minStack = append(this.minStack, x)
//	}
//}
//
//func (this *MinStack) Pop() {
//	if len(this.minStack) > 0 && this.minStack[len(this.minStack)-1] == this.stack[len(this.stack)-1] {
//		this.minStack = this.minStack[:len(this.minStack)-1]
//	}
//	this.stack = this.stack[:len(this.stack)-1]
//}
//
//func (this *MinStack) Top() int {
//	return this.stack[len(this.stack)-1]
//}
//
//func (this *MinStack) GetMin() int {
//	if len(this.minStack) == 0 && len(this.stack) > 0 {
//		this.minStack = append(this.minStack, this.stack[len(this.stack)-1])
//		for i := len(this.stack) - 2; i >= 0; i-- {
//			if this.stack[i] <= this.minStack[len(this.minStack)-1] {
//				this.minStack = append(this.minStack, this.stack[i])
//			}
//		}
//	}
//	return this.minStack[len(this.minStack)-1]
//}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
//leetcode submit region end(Prohibit modification and deletion)
