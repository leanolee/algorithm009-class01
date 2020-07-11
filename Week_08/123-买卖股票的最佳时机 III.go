//给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。 
//
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。 
//
// 注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。 
//
// 示例 1: 
//
// 输入: [3,3,5,0,0,3,1,4]
//输出: 6
//解释: 在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
//     随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。 
//
// 示例 2: 
//
// 输入: [1,2,3,4,5]
//输出: 4
//解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。   
//     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。   
//     因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
// 
//
// 示例 3: 
//
// 输入: [7,6,4,3,1] 
//输出: 0 
//解释: 在这个情况下, 没有交易完成, 所以最大利润为 0。 
// Related Topics 数组 动态规划
package main

import "fmt"

func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	profit := maxProfit(prices)
	fmt.Println(profit)
}

/*
dp：
	股票先买入，才能卖出，所以有：
	第 n 次卖股票的现金 = 第 n 次买入股票的存量 + prices[i]
	第 n 次买入股票的存量 = 第 n-1 次股票卖出后的现金 - prices[i]
*/
//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int) int {
	// dp：
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
	cash, stock := make([][3]int, n), make([][3]int, n)
	stock[0][1], stock[0][2] = -prices[0], -prices[0]
	for i := 1; i < n; i++ {
		for j := 1; j <= 2; j++ {
			cash[i][j] = max(cash[i-1][j], stock[i-1][j]+prices[i])
			stock[i][j] = max(stock[i-1][j], cash[i-1][j-1]-prices[i])
		}
	}
	return cash[n-1][2]
}

//leetcode submit region end(Prohibit modification and deletion)
