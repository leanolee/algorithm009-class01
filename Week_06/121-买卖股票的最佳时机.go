//给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。 
//
// 如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。 
//
// 注意：你不能在买入股票前卖出股票。 
//
// 
//
// 示例 1: 
//
// 输入: [7,1,5,3,6,4]
//输出: 5
//解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
// 
//
// 示例 2: 
//
// 输入: [7,6,4,3,1]
//输出: 0
//解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
// 
// Related Topics 数组 动态规划
package main

import (
	"fmt"
)

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	profit01 := maxProfit01(prices)
	fmt.Println(profit01)
}

//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit01(prices []int) int {
	// dp：
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(prices)
	if n == 0 {
		return 0
	}
	cash, stock := make([]int, n), make([]int, n)
	stock[0] = -prices[0]
	for i := 1; i < n; i++ {
		cash[i] = max(cash[i-1], stock[i-1]+prices[i])
		stock[i] = max(stock[i-1], -prices[i])
	}
	return cash[n-1]

	// dp：
	//min, max := math.MaxInt32, 0
	//for i := 0; i < len(prices); i++ {
	//	if prices[i] < min {
	//		min = prices[i]
	//	}
	//	if prices[i]-min > max {
	//		max = prices[i] - min
	//	}
	//}
	//return max
}

//leetcode submit region end(Prohibit modification and deletion)
