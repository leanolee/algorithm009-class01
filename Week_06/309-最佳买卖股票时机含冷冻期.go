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
	//prices := []int{4, 2, 3, 0, 2}
	prices := []int{2, 1, 4}
	profit04 := maxProfit04(prices)
	fmt.Println(profit04)
}

//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit04(prices []int) int {
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
	dp := make([][2]int, n)
	dp[0][1] = -prices[0]
	pre, curr := 0, 0	// 记录 上上天 和 上天 的现金
	for i := 1; i < n; i++ {
		curr = dp[i-1][0]
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], pre-prices[i])
		pre = curr
	}
	return dp[n-1][0]
}

//leetcode submit region end(Prohibit modification and deletion)
