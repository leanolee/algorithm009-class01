//给定一个非空二维矩阵 matrix 和一个整数 k，找到这个矩阵内部不大于 k 的最大矩形和。 
//
// 示例: 
//
// 输入: matrix = [[1,0,1],[0,-2,3]], k = 2
//输出: 2 
//解释: 矩形区域 [[0, 1], [-2, 3]] 的数值和是 2，且 2 是不超过 k 的最大数字（k = 2）。
// 
//
// 说明： 
//
// 
// 矩阵内的矩形区域面积必须大于 0。 
// 如果行数远大于列数，你将如何解答呢？ 
// 
// Related Topics 队列 二分查找 动态规划
package main

import (
	"fmt"
	"math"
)

func main() {
	matrix := [][]int{{2, 2, -1}}
	//matrix := [][]int{{5, -4, -3, 4}, {-3, -4, 4, 5}, {5, 1, 5, -4}}
	k := 0
	submatrix := maxSumSubmatrix(matrix, k)
	fmt.Println(submatrix)
}

/*
参考：https://leetcode-cn.com/problems/max-sum-of-rectangle-no-larger-than-k/solution/javacong-bao-li-kai-shi-you-hua-pei-tu-pei-zhu-shi/
	https://leetcode-cn.com/problems/max-sum-of-rectangle-no-larger-than-k/solution/bao-li-onyou-hua-er-fen-er-fen-jian-zhi-by-lzh_yve/
	状态转移方程：F(i,j) = F(i-1,j) + F(i,j-1) - F(i-1,j-1) + m(i,j)
暴力法：
*/
//leetcode submit region begin(Prohibit modification and deletion)
func maxSumSubmatrix(matrix [][]int, k int) int {
	// dp：二分+剪枝

	// dp：数组
	max, row, col := math.MinInt32, len(matrix), len(matrix[0])
	for r := 0; r < row; r++ {
		dp := make([]int, col)
		for i := r; i < row; i++ {
			for j := 0; j < col; j++ {
				dp[j] += matrix[i][j]
			}
			//fmt.Println(dp)
			for m := 0; m < col; m++ {
				curr := 0
				for n := m; n < col; n++ {
					curr += dp[n]
					if curr <= k && curr > max {
						max = curr
					}
					if max == k {
						return k
					}
				}
			}
		}
	}
	return max

	// 暴力法：超时，国际站通过
	//max, row, col := math.MinInt32, len(matrix), len(matrix[0])
	//dp := make([][]int, row+1)
	//dp[0] = make([]int, col+1)
	//for r := 1; r <= row; r++ {
	//	for c := 1; c <= col; c++ {
	//		for i := 1; i <= row; i++ {
	//			dp[i] = make([]int, col+1)
	//		}
	//		dp[r][c] = matrix[r-1][c-1]
	//		for i := r; i <= row; i++ {
	//			for j := c; j <= col; j++ {
	//				dp[i][j] = dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1] + matrix[i-1][j-1]
	//				if dp[i][j] == k {
	//					return k
	//				} else if dp[i][j] < k && dp[i][j] > max {
	//					max = dp[i][j]
	//				}
	//			}
	//		}
	//	}
	//}
	//return max
}

//leetcode submit region end(Prohibit modification and deletion)
