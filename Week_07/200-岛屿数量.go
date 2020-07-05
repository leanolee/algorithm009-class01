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
	// 并查集
	if len(grid) == 0 {
		return 0
	}
	count, r, c := 0, len(grid), len(grid[0])
	find := func(parent []int, p int) int {
		for p != parent[p] {
			parent[p], p = parent[parent[p]], parent[parent[p]]
		}
		return p
	}
	union := func(parent []int, p, q int) {
		rootP, rootQ := find(parent, p), find(parent, q)
		if rootP != rootQ {
			parent[rootP] = rootQ
			count--	// 合并一次就减一次
		}
	}
	//arr := []int{0, 3, 4, 2, 0} // test
	//i := find(arr, 1)
	//fmt.Println(i, arr)
	//union(arr, 0, 4)
	dx, dy := [4]int{-1, 0, 0, 1}, [4]int{0, -1, 1, 0}
	parent := make([]int, r*c)
	for i, _ := range parent {
		parent[i] = i // 仍然是必要的步骤
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == '1' {
				count++
				for k := 0; k < 4; k++ {
					nX, nY := i+dx[k], j+dy[k]
					if nX >= 0 && nX < r && nY >= 0 && nY < c && grid[nX][nY] == '1' {
						union(parent, nX*c+nY, i*c+j)
						//count++
					}
				}
			} else {
				parent[i*c+j] = -1
			}
		}
	}
	fmt.Println(parent)
	//count := 0 // count 可以在并查集的两个函数中进行维护：初始为 '1' 的个数，union一次就 count--
	//for i, p := range parent {
	//	if i == p {
	//		count++
	//	}
	//}
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
