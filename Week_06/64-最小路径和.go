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

import "fmt"

func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	sum := minPathSum(grid)
	fmt.Println(sum)
}

//leetcode submit region begin(Prohibit modification and deletion)
func minPathSum(grid [][]int) int {
	minNum := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	r, c := len(grid), len(grid[0])
	dp := make([][]int, r)
	dp[0] = make([]int, c)
	dp[0][0] = grid[0][0]
	for j := 1; j < c; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < r; i++ {
		dp[i] = make([]int, c)
		dp[i][0] = dp[i-1][0] + grid[i][0]
		for j := 1; j < c; j++ {
			dp[i][j] = minNum(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[r-1][c-1]
}

//leetcode submit region end(Prohibit modification and deletion)
