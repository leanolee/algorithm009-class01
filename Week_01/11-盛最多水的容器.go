//给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, 
//ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。 
//
// 说明：你不能倾斜容器，且 n 的值至少为 2。 
//
// 
//
// 
//
// 图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。 
//
// 
//
// 示例： 
//
// 输入：[1,8,6,2,5,4,8,3,7]
//输出：49 
// Related Topics 数组 双指针
package main

import (
	"./MyTools"
	"fmt"
)

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	area := maxArea(height)
	fmt.Println(area)
}
/*
暴力法：O(n^2),O(1)
	1.一层循环：固定一个数
	2.二层循环：遍历数组，取2数之间的最小值作为高，计算max
双指针：O(n),O(1)
	1.循环遍历：指针分别在数组的头、尾，向中推进，直到相遇
	2.计算水量，再取两个指针中值小的向尾/向头推进
*/
//leetcode submit region begin(Prohibit modification and deletion)
func maxArea(height []int) int {
	// 双指针
	max := 0
	for i, j := 0, len(height)-1; i < j; {
		if height[i] < height[j] {
			max = MyTools.Max((j-i)*height[i], max)
			i++
		} else {
			max = MyTools.Max((j-i)*height[j], max)
			j--
		}
	}
	return max

	// 暴力法
	//n := len(height)
	//max := 0
	//for i := 0; i < n-1; i++ {
	//	for j := i + 1; j < n; j++ {
	//		max = MyTools.Max((j-1)*MyTools.Min(height[i], height[j]), max)
	//	}
	//}
	//return max
}

//leetcode submit region end(Prohibit modification and deletion)
