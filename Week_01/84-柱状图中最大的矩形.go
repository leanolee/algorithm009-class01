//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。 
//
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。 
//
// 
//
// 
//
// 以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。 
//
// 
//
// 
//
// 图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。 
//
// 
//
// 示例: 
//
// 输入: [2,1,5,6,2,3]
//输出: 10 
// Related Topics 栈 数组
package main

import (
	"./MyTools"
	"fmt"
)

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	//heights := []int{1, 1}
	//heights := []int{2, 3}
	area := largestRectangleArea(heights)
	fmt.Println(area)
}
/*
暴力法：O(n^3),O(1)
	思路：i 固定一个元素，j 固定另一个元素，寻找 i 到 j 之间的最小元素。距离 * 最小高度 = area
左右边界：O(n^2),O(1)
	思路：固定 0~len-1 之间的元素，寻找其左、右边界。如果是大于当前元素，则继续找，小于则为边界
Stack：O(n),O(n)
	思路：具有最近相关性，可以使用栈来解决
		如果 当前元素 > 栈顶元素，直接压入栈
		如果 当前元素 < 栈顶元素，一直弹出元素，直到当前元素 > 栈顶元素，压入

		可以通过弦曲线来理解
		并且，第一种情况是不确定 右边界；第二种情况是确定了 左边界
	1.创建一个二维数组，第二维为2个元素，分别记录 索引、值
		理解：栈底先放一个 index 为 -1 的元素
	2.第一层循环：遍历元素，压入元素
	3.第二层循环：弹出 大 的元素，计算area：见下
	4.如果栈中还有元素，说明这些元素的 右边界 在遍历完后才确定
		计算栈中剩余元素的area
Stack优化：O(n),O(n)
	变为一维数组
	area = heights[stack[len(stack)-1]] * (i-1-stack[len(stack)-2])
*/
//leetcode submit region begin(Prohibit modification and deletion)
func largestRectangleArea(heights []int) int {
	// Stack：优化
	n := len(heights)
	stack := []int{-1}
	maxArea := 0
	for i := 0; i < n; i++ {
		for len(stack) > 1 && heights[i] < heights[stack[len(stack)-1]] {
			maxArea = MyTools.Max(heights[stack[len(stack)-1]]*(i-1-stack[len(stack)-2]), maxArea)
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for i := len(stack) - 1; i > 0; i-- {
		maxArea = MyTools.Max(heights[stack[i]]*(stack[len(stack)-1]-stack[i-1]), maxArea)
	}
	return maxArea

	// Stack
	//n := len(heights)
	//stack := [][2]int{{-1, 0}}
	//maxArea := 0
	//for i := 0; i < n; i++ {
	//	for len(stack) > 1 && heights[i] < stack[len(stack)-1][1] {
	//		maxArea = MyTools.Max(stack[len(stack)-1][1]*(i-1-stack[len(stack)-2][0]), maxArea)
	//		stack = stack[:len(stack)-1]
	//	}
	//	stack = append(stack, [2]int{i, heights[i]})
	//}
	//for i := len(stack) - 1; i > 0; i-- {
	//	maxArea = MyTools.Max(stack[i][1]*(stack[len(stack)-1][0]-stack[i-1][0]), maxArea)
	//}
	//return maxArea

	// 左右边界
	//n := len(heights)
	//maxArea := 0
	//left, right := 0, 0
	//for i := 0; i < n; i++ {
	//	left, right = i-1, i+1
	//	for left >= 0 {
	//		if heights[i] > heights[left] {
	//			break
	//		}
	//		left--
	//	}
	//	for right < n {
	//		if heights[i] > heights[right] {
	//			break
	//		}
	//		right++
	//	}
	//	maxArea = MyTools.Max(heights[i]*(right-left-1), maxArea)
	//}
	//return maxArea

	// 暴力法
	//n := len(heights)
	//maxArea := 0
	//for i := 0; i < n; i++ {
	//	for j := 0; j < n; j++ {
	//		minHeight := heights[i]
	//		for k := i + 1; k <= j; k++ {
	//			if heights[k] < minHeight {
	//				minHeight = heights[k]
	//			}
	//		}
	//		maxArea = MyTools.Max(minHeight*(j-i+1), maxArea)
	//	}
	//}
	//return maxArea
}

//leetcode submit region end(Prohibit modification and deletion)
