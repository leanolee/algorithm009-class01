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

import "fmt"

func main() {
	//heights := []int{2, 1, 5, 6, 2, 3}
	heights := []int{5, 4, 1, 2}
	area := largestRectangleArea(heights)
	fmt.Println(area)
}

/*
stack：
	stack中记录数组的索引：新增索引对应元素大于等于stack尾，直接添加；否则 for loop ——> pop
	高度：stack尾元素对应的高
	宽度：新增元素 - 1 - 尾元素的上一个元素
*/
//leetcode submit region begin(Prohibit modification and deletion)
func largestRectangleArea(heights []int) int {
	// stack
	stack, maxArea := []int{-1}, 0
	for i, h := range heights {
		for last := stack[len(stack)-1]; last >= 0 && h < heights[last]; last = stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			currArea := heights[last] * (i - 1 - stack[len(stack)-1])
			if currArea > maxArea {
				maxArea = currArea
			}
		}
		stack = append(stack, i)
	}
	for last := stack[len(stack)-1]; last >= 0; last = stack[len(stack)-1] {
		stack = stack[:len(stack)-1]
		currArea := heights[last] * (len(heights) - 1 - stack[len(stack)-1])
		if currArea > maxArea {
			maxArea = currArea
		}
	}
	return maxArea

	// 边界法
}

//leetcode submit region end(Prohibit modification and deletion)
