//给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用原地算法。 
//
// 示例 1: 
//
// 输入: 
//[
//  [1,1,1],
//  [1,0,1],
//  [1,1,1]
//]
//输出: 
//[
//  [1,0,1],
//  [0,0,0],
//  [1,0,1]
//]
// 
//
// 示例 2: 
//
// 输入: 
//[
//  [0,1,2,0],
//  [3,4,5,2],
//  [1,3,1,5]
//]
//输出: 
//[
//  [0,0,0,0],
//  [0,4,5,0],
//  [0,3,1,0]
//] 
//
// 进阶: 
//
// 
// 一个直接的解决方案是使用 O(mn) 的额外空间，但这并不是一个好的解决方案。 
// 一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。 
// 你能想出一个常数空间的解决方案吗？ 
// 
// Related Topics 数组
package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix)
	fmt.Println(matrix)
}

/*
目标：O(1)的空间
暴力：
	1.遍历原始矩阵，如果发现如果某个元素 cell[i][j] 为 0，我们将第 i 行和第 j 列的所有非零元素设成很大的负虚拟值（比如说 -1000000）
		注意，正确的虚拟值取值依赖于问题的约束，任何允许值范围外的数字都可以作为虚拟之
	2.最后，我们便利整个矩阵将所有等于虚拟值（常量在代码中初始化为 MODIFIED）的元素设为 0
记录：
	位运算：
*/
//leetcode submit region begin(Prohibit modification and deletion)
func setZeroes(matrix [][]int) {
	// 暴力：leetcode：O(MN),O(1)
	row, col := len(matrix), len(matrix[0])
	zeroCol := false
	for i := 0; i < row; i++ {
		if matrix[i][0] == 0 { // 记录首列
			zeroCol = true
		}
		for j := 1; j < col; j++ { // 从1开始
			if matrix[i][j] == 0 { // 将需要处理的 行 列 记录到 0行 0列
				matrix[i][0], matrix[0][j] = 0, 0
			}
		}
	}
	for i := 1; i < row; i++ { // 处理非首列首行
		for j := 1; j < col; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if matrix[0][0] == 0 { // 处理首行
		for j := 1; j < col; j++ {
			matrix[0][j] = 0
		}
	}
	if zeroCol { // 处理首列
		for i := 0; i < row; i++ { // 从0开始
			matrix[i][0] = 0
		}
	}

	// DFS：其实也是记录，或者用 负虚拟值
	//row, col := len(matrix), len(matrix[0])
	//var dfs func(x, y int)
	//dfs = func(x, y int) {
	//	for i := x - 1; i >= 0; i-- {
	//		if matrix[i][y] != 0 {
	//			matrix[i][y] = math.MinInt64
	//		}
	//	} // 上
	//	for j := y - 1; j >= 0; j-- {
	//		if matrix[x][j] != 0 {
	//			matrix[x][j] = math.MinInt64
	//		}
	//	} // 左
	//	for j := y + 1; j < col; j++ {
	//		if matrix[x][j] != 0 {
	//			matrix[x][j] = math.MinInt64
	//		}
	//	} // 右
	//	for i := x + 1; i < row; i++ {
	//		if matrix[i][y] != 0 {
	//			matrix[i][y] = math.MinInt64
	//		}
	//	} // 下
	//}
	//for i := 0; i < row; i++ {
	//	for j := 0; j < col; j++ {
	//		if matrix[i][j] == 0 {
	//			dfs(i, j)
	//		}
	//	}
	//}
	//for i := 0; i < row; i++ {
	//	for j := 0; j < col; j++ {
	//		if matrix[i][j] == math.MinInt64 {
	//			matrix[i][j] = 0
	//		}
	//	}
	//}

	// 数组：记录
	//r, c, row, col := make([]bool, len(matrix)), make([]bool, len(matrix[0])), len(matrix), len(matrix[0])
	//for i := 0; i < row; i++ {
	//	for j := 0; j < col; j++ {
	//		if matrix[i][j] == 0 {
	//			r[i], c[j] = true, true
	//		}
	//	}
	//}
	//for i := 0; i < row; i++ {
	//	for j := 0; j < col; j++ {
	//		if matrix[i][j] != 0 && (r[i] || c[j]) {
	//			matrix[i][j] = 0
	//		}
	//	}
	//}
	// 位运算：记录，但是测试样例的 matrix 长度超过了位
	//r, c, row, col := int64(0), int64(0), len(matrix), len(matrix[0])
	//for i := 0; i < row; i++ {
	//	for j := 0; j < col; j++ {
	//		if matrix[i][j] == 0 {
	//			r |= 1 << i
	//			c |= 1 << j
	//		}
	//	}
	//}
	//for i := 0; i < row; i++ {
	//	for j := 0; j < col; j++ {
	//		if matrix[i][j] != 0 && (r&(1<<i) > 0 || c&(1<<j) > 0) {
	//			matrix[i][j] = 0
	//		}
	//	}
	//}
}

//leetcode submit region end(Prohibit modification and deletion)
