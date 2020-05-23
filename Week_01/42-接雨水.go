//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。 
//
// 
//
// 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Mar
//cos 贡献此图。 
//
// 示例: 
//
// 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
//输出: 6 
// Related Topics 栈 数组 双指针
package main

import (
	"fmt"
)

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	//height := []int{4, 2, 3}
	volume := trap(height)
	fmt.Println(volume)
}

/*
暴力法：O(n^2),O(1)
	思路：每一个元素，找它左、右最高的元素
	1.一层循环：固定当前元素
	2.二层循环：分别往左、右找最高元素，再取两者较小值，min
		若 min > height[current]
		sum += min - height[current]
dp：O(n),O(n)
	思路：灵感来自于 leetcode-239
	1.创建两个数组，分别记录：左->右，右->左
		递推的 height[left]、height[right] 两者的最大值 分别与 height[current] 的差值
	2.一次循环：完成上一步
	3.取两个数组中较小的值，累加
Stack：O(n),O(n)
	思路：具有最近相关性，可用栈来解决
	1.创建stack
		新元素若比栈顶元素小，直接入栈
		若比栈顶元素大，则依次弹出，计算
	2.一次循环：(min(stack[top-1], height[i]) - stack[top]) * (i - Idx(top-1) - 1)
		将结果累加
双指针：O(n),O(1)
	思路：来自dp，dp方法中两个数组的数据规律：最大元素的左侧 都小于 最大元素的右侧；反之亦然
	1.左右变量 left right 分别记录，左右递推的最大元素
	2.一次循环：哪边值更小，则哪边往中间推进
		左右分别计算，如：
			height[i] < left，累加差值
			height[i] > left，left = height[i]
*/
//leetcode submit region begin(Prohibit modification and deletion)
func trap(height []int) int {
	// 双指针
	volumes := 0
	left, right := 0, 0
	for i, j := 0, len(height)-1; i < j; {
		if height[i] < height[j] {
			if height[i] < left {
				volumes += left - height[i]
			} else {
				left = height[i]
			}
			i++
		} else {
			if height[j] < right {
				volumes += right - height[j]
			} else {
				right = height[j]
			}
			j--
		}
	}
	return volumes

	// Stack
	//n := len(height)
	//stack := make([]int, 0)
	//volumes := 0
	//for i := 0; i < n; i++ {
	//	for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
	//		shortIndex := stack[len(stack)-1]
	//		stack = stack[:len(stack)-1]
	//		if len(stack) == 0 {
	//			break
	//		}
	//		volumes += (MyTools.Min(height[stack[len(stack)-1]], height[i]) - height[shortIndex]) * (i - stack[len(stack)-1] - 1)
	//	}
	//	stack = append(stack, i)
	//}
	//return volumes

	// dp
	//n := len(height)
	//if n < 3 {
	//	return 0
	//}
	//left := make([]int, n)
	//right := make([]int, n)
	//leftHeiMax, rightHeiMax := 0, 0
	//for i, j := 0, n-1; i < n; {
	//	if height[i] > leftHeiMax {
	//		leftHeiMax = height[i]
	//	}
	//	left[i] = leftHeiMax - height[i]
	//	i++
	//	if height[j] > rightHeiMax {
	//		rightHeiMax = height[j]
	//	}
	//	right[j] = rightHeiMax - height[j]
	//	j--
	//}
	//maxVolume := 0
	//for i := 0; i < n; i++ {
	//	if left[i] < right[i] {
	//		maxVolume += left[i]
	//	} else {
	//		maxVolume += right[i]
	//	}
	//}
	//return maxVolume

	// 暴力法
	//n := len(height)
	//if n < 3 {
	//	return 0
	//}
	//var volumes int
	//for i := 1; i < n-1; i++ {
	//	var maxLeft, maxRight int
	//	// 找左侧最高
	//	for j := i - 1; j >= 0; j-- {
	//		if height[j] > maxLeft {
	//			maxLeft = height[j]
	//		}
	//	}
	//	// 找右侧最高
	//	for k := i + 1; k < n; k++ {
	//		if height[k] > maxRight {
	//			maxRight = height[k]
	//		}
	//	}
	//	volume := MyTools.Min(maxLeft, maxRight) - height[i]
	//	if volume > 0 {
	//		volumes += volume
	//	}
	//}
	//return volumes
}

//leetcode submit region end(Prohibit modification and deletion)
