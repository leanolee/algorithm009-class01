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

import (
	"fmt"
)

func main() {
	//str := "[[\"1\",\"0\",\"1\",\"1\",\"1\"],[\"1\",\"0\",\"1\",\"0\",\"1\"],[\"1\",\"1\",\"1\",\"0\",\"1\"]]"
	//str := "[['B', '1', 'E', '1', 'B'],\n ['B', '1', 'M', '1', 'B'],\n ['B', '1', '1', '1', 'B'],\n ['B', 'B', 'B', 'B', 'B']]"
	//all := strings.ReplaceAll(str, "\"", "'")
	//all = strings.ReplaceAll(all, "[", "{")
	//all = strings.ReplaceAll(all, "]", "}")
	//fmt.Println(all)
	//grid := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}
	grid := [][]byte{{'1', '0', '1', '1', '1'}, {'1', '0', '1', '0', '1'}, {'1', '1', '1', '0', '1'}}
	islands := numIslands(grid)
	fmt.Println(islands)
}

//leetcode submit region begin(Prohibit modification and deletion)
func numIslands(grid [][]byte) int {
	// DFS
	if len(grid) == 0 {
		return 0
	}
	r, c := len(grid), len(grid[0])
	var islandsDFS func(i, j int, from int)
	islandsDFS = func(i, j int, from int) {
		if i < 0 || j < 0 || i >= r || j >= c || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		if from&1 == 0 {
			islandsDFS(i-1, j, 8)
		}
		if from&2 == 0 {
			islandsDFS(i, j-1, 4)
		}
		if from&4 == 0 {
			islandsDFS(i, j+1, 2)
		}
		if from&8 == 0 {
			islandsDFS(i+1, j, 1)
		}
	}
	num := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == '1' {
				num++
				islandsDFS(i, j, 3)
			}
		}
	}
	return num

	// BFS
	//row := len(grid)
	//if grid == nil || row == 0 {
	//	return 0
	//}
	//col := len(grid[0])
	//queue := make([][2]int, 0)
	//numIslands := 0
	//for i := 0; i < row; i++ {
	//	for j := 0; j < col; j++ {
	//		if grid[i][j] == '0' {
	//			continue
	//		}
	//		numIslands++
	//		queue = append(queue, [2]int{i, j})
	//		for len(queue) > 0 {
	//			n := len(queue)
	//			for k := 0; k < n; k++ {
	//				x, y := queue[k][0], queue[k][1]
	//				if x > 0 && grid[x-1][y] == '1' { // 上
	//					grid[x-1][y] = '0'
	//					queue = append(queue, [2]int{x - 1, y})
	//				}
	//				if y > 0 && grid[x][y-1] == '1' { // 左
	//					grid[x][y-1] = '0'
	//					queue = append(queue, [2]int{x, y - 1})
	//				}
	//				if y+1 < col && grid[x][y+1] == '1' { // 右
	//					grid[x][y+1] = '0'
	//					queue = append(queue, [2]int{x, y + 1})
	//				}
	//				if x+1 < row && grid[x+1][y] == '1' { // 下
	//					grid[x+1][y] = '0'
	//					queue = append(queue, [2]int{x + 1, y})
	//				}
	//			}
	//			queue = queue[n:]
	//		}
	//	}
	//}
	//return numIslands

	// DFS
	//row := len(grid)
	//if grid == nil || row == 0 {
	//	return 0
	//}
	//col := len(grid[0])
	//numIslands := 0
	//for i := 0; i < row; i++ {
	//	for j := 0; j < col; j++ {
	//		if grid[i][j] == '1' {
	//			numIslands++
	//			DFS(grid, row, col, i, j, 3) // 起始位置，排除 上 左
	//		}
	//	}
	//}
	//return numIslands

	// 并查集
	numIslands := 0
	// TODO
	return numIslands
}

func DFS(grid [][]byte, row, col, i, j int, path int) {
	if i < 0 || j < 0 || i >= row || j >= col || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	/*
		path 上 左 右 下:=1 2 4 8，如果 (1 2 4 8)&path > 0，则该方向为发起的方向
		来参判断：
			1：排除上
			2：排除左
			4：排除右
			8：排除下
		去参判断：
			1：我是上，去下。你 排除下
			2：我是左，去右。你 排除下
			4：我是右，去左。你 排除下
			8：我是下，去上。你 排除下
		其实根据遍历的特征：可以永远不用往上找（错误，因为有比如：右边先被0隔断，往下绕了再绕上去的情况）
	*/
	if path&1 == 0 {
		DFS(grid, row, col, i-1, j, 8)
	}
	if path&2 == 0 {
		DFS(grid, row, col, i, j-1, 4)
	}
	if path&4 == 0 {
		DFS(grid, row, col, i, j+1, 2)
	}
	if path&8 == 0 {
		DFS(grid, row, col, i+1, j, 1)
	}
	//DFS(grid, row, col, i-1, j, 8)
	//DFS(grid, row, col, i, j+1, 2)
	//DFS(grid, row, col, i, j-1, 4)
	//DFS(grid, row, col, i+1, j, 1)
}

//leetcode submit region end(Prohibit modification and deletion)
