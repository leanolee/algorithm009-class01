//给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。 
//
// 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）: 
//
// 
// 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。 
// 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。 
// 
//
// 示例: 
//
// 输入: [1,2,3,0,2]
//输出: 3 
//解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出] 
// Related Topics 动态规划
package main

import "fmt"

func main() {
	prices := []int{1, 2, 3, 0, 2}
	profit := maxProfit(prices)
	fmt.Println(profit)
}

/*
dp：
	包含冷冻期 n 天：
	将前 n 天的现金余额保存下来，以便后续买入
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
	cash, stock, last, pre := make([]int, n), make([]int, n), 0, 0
	stock[0] = -prices[0]
	for i := 1; i < n; i++ {
		cash[i] = max(cash[i-1], stock[i-1]+prices[i])
		stock[i] = max(stock[i-1], last-prices[i])
		last, pre = pre, cash[i]
	}
	return cash[n-1]
}

//leetcode submit region end(Prohibit modification and deletion)
