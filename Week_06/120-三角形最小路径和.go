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

import (
	"fmt"
	"math"
)

func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	total := minimumTotal(triangle)
	fmt.Println(total)
}

/*
国际站参考：https://leetcode.com/problems/triangle/discuss/38735/Python-easy-to-understand-solutions-(top-down-bottom-up)

*/
//leetcode submit region begin(Prohibit modification and deletion)
func minimumTotal(triangle [][]int) int {
	// dp：自底向上？

	// dp：O(n)空间，自顶向下？
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	r, c := len(triangle), len(triangle[len(triangle)-1])
	dp := make([]int, c+1)
	dp[1] = triangle[0][0]   // 初始化第 0 行
	for i := 1; i < r; i++ { // 从第 1 行开始
		w := len(triangle[i])
		dp[w] = triangle[i][w-1] + dp[w-1] // 先算尾
		for j := w - 1; j > 1; j-- {       // 再算中
			dp[j] = min(dp[j], dp[j-1]) + triangle[i][j-1]
		}
		dp[1] = triangle[i][0] + dp[1] // 后算头
	}
	minTotal := math.MaxInt64
	for i := 1; i <= c; i++ {
		minTotal = min(dp[i], minTotal)
	}
	return minTotal

	// 递归：递归都写成dp了？参见：https://leetcode-cn.com/problems/triangle/solution/di-gui-ji-yi-hua-sou-suo-zai-dao-dp-by-crsm/
	//minimumTotalRecursion()
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func minimumTotalRecursion(triangle [][]int, memo *[]int, r int, idx int) {
}

//leetcode submit region end(Prohibit modification and deletion)
