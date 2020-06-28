//给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。 
//
// 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。 
//
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。 
//
// 
//
// 示例 1: 
//
// 输入: [7,1,5,3,6,4]
//输出: 7
//解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
//     随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
// 
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
//解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。 
//
// 
//
// 提示： 
//
// 
// 1 <= prices.length <= 3 * 10 ^ 4 
// 0 <= prices[i] <= 10 ^ 4 
// 
// Related Topics 贪心算法 数组
package main

import "fmt"

func main() {
	//prices := []int{3, 6, 7, 1, 5, 0, 2, 8, 4} // 16
	prices := []int{3, 2, 5, 7} // 16
	profit := maxProfit02(prices)
	fmt.Println(profit)
}

//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit02(prices []int) int {
	// dp：
	/*
		1.cash[i] = Max(cash[i-1], stock[i-1]+prices[i])
			今天最大现金 = max(昨天最大现金, 昨天屯的股票 + 今天的价格)
		2.stock[i] = Max(stock[i-1], cash[i-1]-prices[i])
			今天屯的股票 = max(昨天屯的股票, 昨天最大现金 - 今天的价格)
	*/
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(prices)
	cash, stock := make([]int, n), make([]int, n)
	cash[0], stock[0] = 0, -prices[0]
	for i, j := 1, 0; i < n; i, j = i+1, i {
		cash[i] = max(cash[j], stock[j]+prices[i])
		stock[i] = max(stock[j], cash[j]-prices[i])
	}
	fmt.Println(cash)
	fmt.Println(stock)
	return cash[n-1]

	// 峰谷法
	//max := 0
	//for i := 1; i < len(prices); i++ {
	//	if prices[i] > prices[i-1] {
	//		max += prices[i] - prices[i-1]
	//	}
	//}
	//return max
}

//leetcode submit region end(Prohibit modification and deletion)
