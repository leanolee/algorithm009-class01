//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。 
//
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。 
//
// 问总共有多少条不同的路径？ 
//
// 
//
// 例如，上图是一个7 x 3 的网格。有多少可能的路径？ 
//
// 
//
// 示例 1: 
//
// 输入: m = 3, n = 2
//输出: 3
//解释:
//从左上角开始，总共有 3 条路径可以到达右下角。
//1. 向右 -> 向右 -> 向下
//2. 向右 -> 向下 -> 向右
//3. 向下 -> 向右 -> 向右
// 
//
// 示例 2: 
//
// 输入: m = 7, n = 3
//输出: 28 
//
// 
//
// 提示： 
//
// 
// 1 <= m, n <= 100 
// 题目数据保证答案小于等于 2 * 10 ^ 9 
// 
// Related Topics 数组 动态规划
package main

import "fmt"

func main() {
	m, n := 51, 10
	paths := uniquePaths(m, n)
	fmt.Println(paths)
}

//leetcode submit region begin(Prohibit modification and deletion)
func uniquePaths(m int, n int) int {
	// dp：一维
	//dp := make([]int, n)
	//for i := 0; i < m; i++ {
	//	dp[0] = 1
	//	for j := 1; j < n; j++ {
	//		dp[j] += dp[j-1]
	//	}
	//}
	//return dp[n-1]

	// dp：二维
	//dp := make([][]int, m+1)
	//for i := 0; i <= m; i++ {
	//	dp[i] = make([]int, n+1)
	//}
	//dp[0][1] = 1
	//for i := 1; i <= m; i++ {
	//	for j := 1; j <= n; j++ {
	//		dp[i][j] = dp[i-1][j] + dp[i][j-1]
	//	}
	//}
	//return dp[m][n]

	// 递归：记忆化
	memo := make(map[int]int)
	count := memoRecursion(m, n, memo)
	fmt.Println(len(memo))
	return count

	// 递归：超时
	//if m == 1 || n == 1 {
	//	return 1
	//}
	//return uniquePaths(m-1, n) + uniquePaths(m, n-1)
}

func memoRecursion(m int, n int, memo map[int]int) int {
	if m == 1 || n == 1 {
		return 1
	}
	idx := m*100 + n - 1
	if memo[idx] == 0 {
		memo[idx] = memoRecursion(m-1, n, memo) + memoRecursion(m, n-1, memo)
	}
	return memo[idx]
}

//leetcode submit region end(Prohibit modification and deletion)
