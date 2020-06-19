//给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。 
//
// 相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。 
//
// 
//
// 例如，给定三角形： 
//
// [
//     [2],
//    [3,4],
//   [6,5,7],
//  [4,1,8,3]
//]
// 
//
// 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。 
//
// 
//
// 说明： 
//
// 如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。 
// Related Topics 数组 动态规划
package main

import "fmt"

func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	total := minimumTotal(triangle)
	fmt.Println(total)
}

//leetcode submit region begin(Prohibit modification and deletion)
func minimumTotal(triangle [][]int) int {
	// dp：正序
	minNum := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	dpLast := make([]int, len(triangle[len(triangle)-1]))
	dpCurr := make([]int, len(triangle[len(triangle)-1]))
	dpLast[0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		n := len(triangle[i]) - 1
		dpCurr[0], dpCurr[n] = triangle[i][0]+dpLast[0], triangle[i][n]+dpLast[n-1]
		for j := 1; j < n; j++ {
			dpCurr[j] = minNum(dpLast[j-1], dpLast[j]) + triangle[i][j]
		}
		copy(dpLast, dpCurr)
	}
	min := dpLast[0]
	for i := 1; i < len(dpLast); i++ {
		if dpLast[i] < min {
			min = dpLast[i]
		}
	}
	return min
}

//leetcode submit region end(Prohibit modification and deletion)
