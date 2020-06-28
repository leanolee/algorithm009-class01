//给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。 
//
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。 
//
// 注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。 
//
// 示例 1: 
//
// 输入: [2,4,1], k = 2
//输出: 2
//解释: 在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
// 
//
// 示例 2: 
//
// 输入: [3,2,6,5,0,3], k = 2
//输出: 7
//解释: 在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
//     随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3 
//。
// 
// Related Topics 动态规划
package main

import (
	"fmt"
	"math"
)

func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4, 7, 2, 9, 6, 8} // 18
	k := 4
	//prices := []int{3, 2, 6, 5, 0, 3}
	//k := 2
	profit05 := maxProfit05(k, prices)
	fmt.Println(profit05)
}

//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit05(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	//dp := make([][][2]int, n)	// out of memory
	//dp[0] = make([][2]int, k+1)
	//for i := 1; i <= k; i++ {
	//	dp[0][i][1] = -prices[0]
	//}
	//for i := 1; i < n; i++ {
	//	dp[i] = make([][2]int, k+1)
	//	for j := 1; j <= k; j++ {
	//		dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
	//		dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
	//	}
	//}
	//return dp[n-1][k][0]
	if k > n>>1 { // 重点 1
		max := 0
		for i := 1; i < n; i++ {
			if prices[i] > prices[i-1] {
				max += prices[i] - prices[i-1]
			}
		}
		return max
	}
	dp0, dp1 := make([]int, k+1), make([]int, k+1)
	for i := 0; i <= k; i++ {
		dp1[i] = math.MinInt32 // 重点 2
	}
	for i := 0; i < n; i++ {
		for j := 1; j <= k; j++ {
			dp0[j] = max(dp0[j], dp1[j]+prices[i])
			dp1[j] = max(dp1[j], dp0[j-1]-prices[i])
		}
	}
	return dp0[k]
	/*
		[[[0 0] [0 -3] [0 -3] [0 -3] [0 -3]]
		[[0 0] [0 -3] [0 -3] [0 -3] [0 -3]]
		[[0 0] [2 -3] [2 -3] [2 -3] [2 -3]]
		[[0 0] [2 0] [2 2] [2 2] [2 2]]
		[[0 0] [2 0] [2 2] [2 2] [2 2]]
		[[0 0] [3 0] [5 2] [5 2] [5 2]]
		[[0 0] [3 0] [5 2] [5 4] [5 4]]
		[[0 0] [4 0] [6 2] [8 4] [8 4]]
		[[0 0] [7 0] [9 2] [11 4] [11 4]]
		[[0 0] [7 0] [9 5] [11 7] [11 9]]
		[[0 0] [9 0] [14 5] [16 7] [18 9]]
		[[0 0] [9 0] [14 5] [16 8] [18 10]]
		[[0 0] [9 0] [14 5] [16 8] [18 10]]]
	*/
}

//leetcode submit region end(Prohibit modification and deletion)
