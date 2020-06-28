//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。 
//
// 岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。 
//
// 此外，你可以假设该网格的四条边均被水包围。 
//
// 
//
// 示例 1: 
//
// 输入:
//11110
//11010
//11000
//00000
//输出: 1
// 
//
// 示例 2: 
//
// 输入:
//11000
//11000
//00100
//00011
//输出: 3
//解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。
// 
// Related Topics 深度优先搜索 广度优先搜索 并查集
package main

import "fmt"

func main() {
	grid := [][]byte{{'1', '0', '1', '1', '1'}, {'1', '0', '1', '0', '1'}, {'1', '1', '1', '0', '1'}}
	islands := numIslands(grid)
	fmt.Println(islands)
}

//leetcode submit region begin(Prohibit modification and deletion)
func numIslands(grid [][]byte) int {
	// dfs
	if len(grid) == 0 {
		return 0
	}
	count, r, c := 0, len(grid), len(grid[0] )
	var dfs func(i, j int, direct int)
	dfs = func(i, j int, direct int) {
		if i < 0 || i >= r || j < 0 || j >= c || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		if direct&1 == 0 {
			dfs(i-1, j, 8)
		}
		if direct&2 == 0 {
			dfs(i, j-1, 4)
		}
		if direct&4 == 0 {
			dfs(i, j+1, 2)
		}
		if direct&8 == 0 {
			dfs(i+1, j, 1)
		}
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == '1' {
				dfs(i, j, 3)
				count++
			}
		}
	}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
