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
	//prices := []int{7, 6, 4, 3, 1}
	prices := []int{7, 1, 5, 3, 6, 4}
	//prices := []int{1, 5}
	profit := maxProfit(prices)
	fmt.Println(profit)
}

/*
峰谷法：
	谷1 + 峰1 + 谷2 + 峰2：试图跳过 峰1 和 谷2，得到的差值是 小于 两次买卖的
*/
//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int) int {
	// 峰谷法：优化
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit

	// 峰谷法
	//n := len(prices) - 1
	//valley, peak := 0, 0 // 表示 谷 和 峰
	//maxProfit := 0
	//i := 0
	//for i < n {
	//	for i < n && prices[i] >= prices[i+1] {
	//		i++ // 找到 谷
	//	}
	//	valley = i
	//	for i < n && prices[i] <= prices[i+1] {
	//		i++ // 找到 峰
	//	}
	//	peak = i
	//	maxProfit += prices[peak] - prices[valley]
	//}
	//return maxProfit

	// 暴力法：超时
	//return maxProfitRecursion(prices, 0)
}

func maxProfitRecursion(prices []int, begin int) int {
	max := 0
	for i := begin; i < len(prices)-1; i++ {
		maxMoney := 0
		for j := i + 1; j < len(prices); j++ {
			if prices[i] < prices[j] {
				money := prices[j] - prices[i] + maxProfitRecursion(prices, j+1)
				if money > maxMoney {
					maxMoney = money
				}
			}
		}
		if maxMoney > max {
			max = maxMoney
		}
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
