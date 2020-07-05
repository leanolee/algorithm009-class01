//在一个 N × N 的方形网格中，每个单元格有两种状态：空（0）或者阻塞（1）。 
//
// 一条从左上角到右下角、长度为 k 的畅通路径，由满足下述条件的单元格 C_1, C_2, ..., C_k 组成： 
//
// 
// 相邻单元格 C_i 和 C_{i+1} 在八个方向之一上连通（此时，C_i 和 C_{i+1} 不同且共享边或角） 
// C_1 位于 (0, 0)（即，值为 grid[0][0]） 
// C_k 位于 (N-1, N-1)（即，值为 grid[N-1][N-1]） 
// 如果 C_i 位于 (r, c)，则 grid[r][c] 为空（即，grid[r][c] == 0） 
// 
//
// 返回这条从左上角到右下角的最短畅通路径的长度。如果不存在这样的路径，返回 -1 。 
//
// 
//
// 示例 1： 
//
// 输入：[[0,1],[1,0]]
//
//输出：2
//
// 
//
// 示例 2： 
//
// 输入：[[0,0,0],[1,1,0],[1,1,0]]
//
//输出：4
//
// 
//
// 
//
// 提示： 
//
// 
// 1 <= grid.length == grid[0].length <= 100 
// grid[i][j] 为 0 或 1 
// 
// Related Topics 广度优先搜索
package main

import (
	"container/list"
	"fmt"
)

func main() {
	grid := [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{1, 1, 0}}
	matrix := shortestPathBinaryMatrix(grid)
	fmt.Println(matrix)
}

//leetcode submit region begin(Prohibit modification and deletion)
func shortestPathBinaryMatrix(grid [][]int) int {
	// A*

	// BFS + dp
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}
	dp, queue, dx, dy := make([][]int, n), list.New(), [8]int{-1, -1, -1, 0, 0, 1, 1, 1}, [8]int{-1, 0, 1, -1, 1, -1, 0, 1}
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	queue.PushBack([2]int{0, 0})
	for queue.Len() > 0 {
		curr := queue.Remove(queue.Front()).([2]int)
		i, j := curr[0], curr[1]
		for k := 0; k < 8; k++ {
			nI, nJ := i+dx[k], j+dy[k]
			if nI < 0 || nI >= n || nJ < 0 || nJ >= n || grid[nI][nJ] == 1 || dp[nI][nJ] > 0 {
				continue
			}
			dp[nI][nJ] = dp[i][j] + 1
			queue.PushBack([2]int{nI, nJ})
		}
	}
	if dp[n-1][n-1] == 0 {
		return -1
	}
	return dp[n-1][n-1]
}

//leetcode submit region end(Prohibit modification and deletion)
