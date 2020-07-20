//给定一个仅包含 0 和 1 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。 
//
// 示例: 
//
// 输入:
//[
//  ["1","0","1","0","0"],
//  ["1","0","1","1","1"],
//  ["1","1","1","1","1"],
//  ["1","0","0","1","0"]
//]
//输出: 6 
// Related Topics 栈 数组 哈希表 动态规划
package main

import "fmt"

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'}}
	rectangle := maximalRectangle(matrix)
	fmt.Println(rectangle)
}

/*
暴力法：
dp：O(m * n^2)
	dp[i][j]：记录第 i 行的最大宽度
		width = min(width, dp[k][j])：k = 0~i
		area = width * (i-k+1)
柱状图：栈 O(mn)
	思路来自 lc-84
	先求出每一行的最大高度
	再求每一行的最大面积
dp：每个点的最大高度

*/
//leetcode submit region begin(Prohibit modification and deletion)
func maximalRectangle(matrix [][]byte) int {
	// dp：每个点的最大高度

	// 柱状图：栈
	if len(matrix) == 0 {
		return 0
	}
	maxArea, m, n := 0, len(matrix), len(matrix[0])
	hs, stack := make([]int, n), []int{-1}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				hs[j]++
			} else {
				hs[j] = 0
			}
		}
		maxArea = maxVal(maxArea, largestArea(hs, stack[:1]))
	}
	return maxArea

	// dp：
	//if len(matrix) == 0 {
	//	return 0
	//}
	//m, n := len(matrix), len(matrix[0])
	//dp, maxArea := make([][]int, m), 0
	//for i := 0; i < m; i++ {
	//	dp[i] = make([]int, n)
	//}
	//for i := 0; i < m; i++ {
	//	for j := 0; j < n; j++ {
	//		if matrix[i][j] == '1' {
	//			if j == 0 {
	//				dp[i][j] = 1
	//			} else {
	//				dp[i][j] = dp[i][j-1] + 1
	//			}
	//		}
	//		width := dp[i][j]
	//		for k := i; k >= 0; k-- {
	//			width = minVal(width, dp[k][j])
	//			maxArea = maxVal(maxArea, width*(i-k+1))
	//		}
	//	}
	//}
	//return maxArea
}

func largestArea(hs []int, stack []int) int {
	maxArea := 0
	for i, h := range hs {
		for last := stack[len(stack)-1]; last >= 0 && h < hs[last]; last = stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			area := hs[last] * (i - 1 - stack[len(stack)-1])
			if area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, i)
	}
	for last := stack[len(stack)-1]; last >= 0; last = stack[len(stack)-1] {
		stack = stack[:len(stack)-1]
		area := hs[last] * (len(hs) - 1 - stack[len(stack)-1])
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
func minVal(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
