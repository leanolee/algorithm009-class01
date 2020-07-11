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

import "fmt"

func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	k := 3
	profit := maxProfit(k, prices)
	fmt.Println(profit)
}

/*
dp：
	股票先买入，才能卖出，所以有：
	第 n 次卖股票的现金 = 第 n 次买入股票的存量 + prices[i]
	第 n 次买入股票的存量 = 第 n-1 次股票卖出后的现金 - prices[i]
*/
//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	if k >= n>>1 {
		maxCash := 0
		for i := 1; i < len(prices); i++ {
			if prices[i] > prices[i-1] {
				maxCash += prices[i] - prices[i-1]
			}
		}
		return maxCash
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	cash, stock := make([][]int, n), make([][]int, n)
	for i := 0; i < n; i++ {
		cash[i] = make([]int, k+1)
		stock[i] = make([]int, k+1)
	}
	for i := 1; i <= k; i++ {
		stock[0][i] = -prices[0]
	}
	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			cash[i][j] = max(cash[i-1][j], stock[i-1][j]+prices[i])
			stock[i][j] = max(stock[i-1][j], cash[i-1][j-1]-prices[i])
		}
	}
	return cash[n-1][k]
}

//leetcode submit region end(Prohibit modification and deletion)
