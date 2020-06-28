//在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。 
//
// 示例: 
//
// 输入: 
//
//1 0 1 0 0
//1 0 1 1 1
//1 1 1 1 1
//1 0 0 1 0
//
//输出: 4 
// Related Topics 动态规划
package main

import (
	"fmt"
	"strings"
)

func main() {
	//s := "[[\"1\",\"0\",\"1\",\"0\",\"0\"],[\"1\",\"0\",\"1\",\"1\",\"1\"],[\"1\",\"1\",\"1\",\"1\",\"1\"],[\"1\",\"0\",\"0\",\"1\",\"0\"]]"
	s := "[\"A\",\"A\",\"A\",\"A\",\"A\",\"A\",\"B\",\"C\",\"D\",\"E\",\"F\",\"G\"]"
	s = strings.ReplaceAll(s, "[", "{")
	s = strings.ReplaceAll(s, "]", "}")
	s = strings.ReplaceAll(s, "\"", "'")
	fmt.Println(s)

	matrix := [][]uint8{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}
	//matrix := [][]byte{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}
	square := maximalSquare(matrix)
	fmt.Println(square)
}

//leetcode submit region begin(Prohibit modification and deletion)
func maximalSquare(matrix [][]uint8) int {
	// dp：
	r := len(matrix)
	if r == 0 {
		return 0
	}
	c := len(matrix[0])
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	dp := make([][]int, r+1)
	dp[0] = make([]int, c+1)
	max := 0
	for i := 1; i <= r; i++ {
		dp[i] = make([]int, c+1)
		for j := 1; j <= c; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
				if dp[i][j] > max {
					max = dp[i][j]
				}
			}
		}
	}
	return max * max

	// 暴力法：
	//if len(matrix) == 0 {
	//	return 0
	//}
	//maxSide, side := 0, 0
	//r, c := len(matrix), len(matrix[0])
	//var getSide func(i, j, m, n int, side int) int
	//getSide = func(i, j, m, n int, side int) int {
	//	if m >= r || n >= c {
	//		return side
	//	}
	//	for col := j; col <= n; col++ {
	//		if matrix[m][col] == '0' {
	//			return side
	//		}
	//	}
	//	for row := i; row <= m; row++ {
	//		if matrix[row][n] == '0' {
	//			return side
	//		}
	//	}
	//	return getSide(i, j, m+1, n+1, side+1)
	//}
	//for i := 0; i < r; i++ {
	//	for j := 0; j < c; j++ {
	//		if matrix[i][j] == '0' {
	//			continue
	//		}
	//		side = getSide(i, j, i+1, j+1, 1)
	//		if side > maxSide {
	//			maxSide = side
	//		}
	//	}
	//}
	//return maxSide * maxSide
}

//leetcode submit region end(Prohibit modification and deletion)
