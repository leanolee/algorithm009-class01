//在二维网格 grid 上，有 4 种类型的方格： 
//
// 
// 1 表示起始方格。且只有一个起始方格。 
// 2 表示结束方格，且只有一个结束方格。 
// 0 表示我们可以走过的空方格。 
// -1 表示我们无法跨越的障碍。 
// 
//
// 返回在四个方向（上、下、左、右）上行走时，从起始方格到结束方格的不同路径的数目，每一个无障碍方格都要通过一次。 
//
// 
//
// 示例 1： 
//
// 输入：[[1,0,0,0],[0,0,0,0],[0,0,2,-1]]
//输出：2
//解释：我们有以下两条路径：
//1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2)
//2. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2) 
//
// 示例 2： 
//
// 输入：[[1,0,0,0],[0,0,0,0],[0,0,0,2]]
//输出：4
//解释：我们有以下四条路径： 
//1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2),(2,3)
//2. (0,0),(0,1),(1,1),(1,0),(2,0),(2,1),(2,2),(1,2),(0,2),(0,3),(1,3),(2,3)
//3. (0,0),(1,0),(2,0),(2,1),(2,2),(1,2),(1,1),(0,1),(0,2),(0,3),(1,3),(2,3)
//4. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2),(2,3) 
//
// 示例 3： 
//
// 输入：[[0,1],[2,0]]
//输出：0
//解释：
//没有一条路能完全穿过每一个空的方格一次。
//请注意，起始和结束方格可以位于网格中的任意位置。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= grid.length * grid[0].length <= 20 
// 
// Related Topics 深度优先搜索 回溯算法
package main

import "fmt"

func main() {
	grid := [][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1}}
	//grid := [][]int{{0, 1}, {2, 0}}
	iii := uniquePathsIII(grid)
	fmt.Println(iii)
}

//leetcode submit region begin(Prohibit modification and deletion)
func uniquePathsIII(grid [][]int) int {
	// dp：hash
	row, col, mark := len(grid), len(grid[0]), 0
	getMark := func(r, c int) int {
		return 1 << (r*col + c)
	}
	memo := make([][]map[int]int, row)
	rBegin, cBegin, rTo, cTo := 0, 0, 0, 0
	for i := 0; i < row; i++ {
		memo[i] = make([]map[int]int, col)
		for j := 0; j < col; j++ {
			memo[i][j] = make(map[int]int)
			switch grid[i][j] {
			case -1:
			case 1:
				rBegin, cBegin = i, j
			case 2:
				rTo, cTo = i, j
				fallthrough
			case 0:
				mark |= getMark(i, j)
			}
		}
	}
	dr, dc := []int{-1, 0, 0, 1}, []int{0, -1, 1, 0}
	var dp func(r, c int, m int) int
	dp = func(r, c int, m int) int {
		if v, ok := memo[r][c][m]; ok {
			return v
		}
		if r == rTo && c == cTo {
			if m == 0 {
				return 1
			}
			return 0
		}
		var ans int
		for i := 0; i < 4; i++ {
			nr, nc := r+dr[i], c+dc[i]
			if nr >= 0 && nr < row && nc >= 0 && nc < col && (m&getMark(nr, nc)) != 0 {
				ans += dp(nr, nc, m^getMark(nr, nc))
			}
		}
		memo[r][c][m] = ans
		return ans
	}
	//return dp(rBegin, cBegin, mark)
	res := dp(rBegin, cBegin, mark)
	fmt.Println(memo)
	return res

	// dp：
	//row, col, mark := len(grid), len(grid[0]), 0
	//getMark := func(r, c int) int {
	//	return 1 << (r*col + c)
	//}
	//memo := make([][][]int, row)
	//rBegin, cBegin, rTo, cTo := 0, 0, 0, 0
	//for i := 0; i < row; i++ {
	//	memo[i] = make([][]int, col)
	//	for j := 0; j < col; j++ {
	//		memo[i][j] = make([]int, 1<<(row*col))
	//		for k := 0; k < 1<<(row*col); k++ {
	//			memo[i][j][k] = -1
	//		}
	//		switch grid[i][j] {
	//		case -1:
	//		case 1:
	//			rBegin, cBegin = i, j
	//			//mark = getMark(i, j, mark)
	//		case 2:
	//			rTo, cTo = i, j
	//			fallthrough
	//		case 0:
	//			mark |= getMark(i, j)
	//		}
	//	}
	//}
	//dr, dc := []int{-1, 0, 0, 1}, []int{0, -1, 1, 0}
	//var dp func(r, c int, m int) int
	//dp = func(r, c int, m int) int {
	//	if memo[r][c][m] >= 0 {
	//		return memo[r][c][m]
	//	}
	//	if r == rTo && c == cTo {
	//		if m == 0 {
	//			return 1
	//		}
	//		return 0
	//	}
	//	var ans int
	//	for i := 0; i < 4; i++ {
	//		nr, nc := r+dr[i], c+dc[i]
	//		if nr >= 0 && nr < row && nc >= 0 && nc < col && (m&getMark(nr, nc)) != 0 {
	//			//if nr >= 0 && nr < row && nc >= 0 && nc < col {
	//			//	if (m & getMark(nr, nc)) != 0 {
	//			ans += dp(nr, nc, m^getMark(nr, nc))
	//			//}
	//		}
	//	}
	//	memo[r][c][m] = ans
	//	return ans
	//}
	//return dp(rBegin, cBegin, mark)

	// DFS：速度最快
	//row, col := len(grid), len(grid[0])
	//count, paths, r, c := 0, row*col, 0, 0
	//for i := 0; i < row; i++ {
	//	for j := 0; j < col; j++ {
	//		if grid[i][j] == -1 {
	//			paths--
	//			continue
	//		}
	//		if grid[i][j] == 1 {
	//			r, c = i, j
	//		}
	//	}
	//}
	//uniquePathsIIIRecursion(grid, r, c, &count, paths)
	//return count
}

func uniquePathsIIIRecursion(grid [][]int, r, c int, count *int, paths int) {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] == -1 || grid[r][c] == 2 && paths > 1 {
		return
	}
	if paths == 1 && grid[r][c] == 2 {
		*count++
		return
	}
	grid[r][c] = -1
	uniquePathsIIIRecursion(grid, r-1, c, count, paths-1)
	uniquePathsIIIRecursion(grid, r, c-1, count, paths-1)
	uniquePathsIIIRecursion(grid, r, c+1, count, paths-1)
	uniquePathsIIIRecursion(grid, r+1, c, count, paths-1)
	grid[r][c] = 0
}

//leetcode submit region end(Prohibit modification and deletion)
