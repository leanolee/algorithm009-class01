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

import "fmt"

func main() {
	//grid := [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}
	grid := [][]int{{0, 1, 0, 0, 0}, {0, 1, 0, 0, 0}, {0, 0, 0, 0, 1}, {0, 1, 1, 1, 0}, {0, 1, 0, 0, 0}}
	matrix := shortestPathBinaryMatrix(grid)
	fmt.Println(matrix)
}

//leetcode submit region begin(Prohibit modification and deletion)
func shortestPathBinaryMatrix(grid [][]int) int {
	// DP：DFS -> F(i,j) = min(F(i-1,j-1),F(i-1,j),F(i-1,j+1),F(i,j-1)...)+1：行不通
	// BFS
	r, c := len(grid), len(grid[0])
	if grid[0][0] == 1 || grid[r-1][c-1] == 1 {
		return -1
	}
	dp := make([][]int, r)
	for i := 0; i < r; i++ {
		dp[i] = make([]int, c)
	}
	dx, dy := [8]int{-1, -1, -1, 0, 0, 1, 1, 1}, [8]int{-1, 0, 1, -1, 1, -1, 0, 1}
	dp[0][0] = 1
	queue := []*Point{{0, 0}}
	//for len(queue) > 0 && dp[r-1][c-1] == 0 {
	//	curr := queue[0]
	//	for i := 0; i < 8; i++ {
	//		nX, nY := curr.x+dx[i], curr.y+dy[i]
	//		if nX >= 0 && nX < r && nY >= 0 && nY < c && grid[nX][nY] == 0 && dp[nX][nY] == 0 {
	//			dp[nX][nY] = dp[curr.x][curr.y] + 1
	//			//if nX == eX && nY == eY { // 找到答案
	//			//	return dp[nX][nY]
	//			//}
	//			queue = append(queue, &Point{nX, nY})
	//		}
	//	}
	//	queue = queue[1:]
	//}
	for idx := 0; idx < len(queue) && dp[r-1][c-1] == 0; idx++ {
		curr := queue[idx]
		for i := 0; i < 8; i++ {
			nX, nY := curr.x+dx[i], curr.y+dy[i]
			if nX >= 0 && nX < r && nY >= 0 && nY < c && grid[nX][nY] == 0 && dp[nX][nY] == 0 {
				dp[nX][nY] = dp[curr.x][curr.y] + 1
				//if nX == eX && nY == eY { // 找到答案
				//	return dp[nX][nY]
				//}
				queue = append(queue, &Point{nX, nY})
			}
		}
	}
	if dp[r-1][c-1] == 0 {
		return -1
	}
	return dp[r-1][c-1]
}

type Point struct {
	x int
	y int
}

//leetcode submit region end(Prohibit modification and deletion)
