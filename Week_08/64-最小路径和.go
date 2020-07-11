//给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。 
//
// 说明：每次只能向下或者向右移动一步。 
//
// 示例: 
//
// 输入:
//[
//  [1,3,1],
//  [1,5,1],
//  [4,2,1]
//]
//输出: 7
//解释: 因为路径 1→3→1→1→1 的总和最小。
// 
// Related Topics 数组 动态规划
package main

import (
	"fmt"
	"math"
)

func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1}}
	sum := minPathSum(grid)
	fmt.Println(sum)
}

//leetcode submit region begin(Prohibit modification and deletion)
func minPathSum(grid [][]int) int {
	// dp：
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	m, n := len(grid), len(grid[0])
	dp := make([]int, n+1)
	for j := 1; j <= n; j++ {
		dp[j] = grid[0][j-1] + dp[j-1]
	}
	dp[0] = math.MaxInt64
	for i := 1; i < m; i++ {
		for j := 1; j <= n; j++ {
			dp[j] = grid[i][j-1] + min(dp[j-1], dp[j])
		}
	}
	return dp[n]
}

//leetcode submit region end(Prohibit modification and deletion)
