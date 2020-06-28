//给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回
// -1。 
//
// 
//
// 示例 1: 
//
// 输入: coins = [1, 2, 5], amount = 11
//输出: 3 
//解释: 11 = 5 + 5 + 1 
//
// 示例 2: 
//
// 输入: coins = [2], amount = 3
//输出: -1 
//
// 
//
// 说明: 
//你可以认为每种硬币的数量是无限的。 
// Related Topics 动态规划
package main

import "fmt"

func main() {
	coins := []int{1, 2, 5}
	amount := 100
	change := coinChange(coins, amount)
	fmt.Println(change)
}

//leetcode submit region begin(Prohibit modification and deletion)
func coinChange(coins []int, amount int) int {
	// dp：
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				dp[i] = min(dp[i], dp[i-coins[j]]+1) // 多个值中取 最小值，就这样写
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]

	// 递归：BFS

	// 递归：DFS，超时，则可以记忆化来优化
	//min := amount + 1
	//coinChangeRecursion(coins, amount, 0, &min)
	//if min == amount+1 {
	//	return -1
	//}
	//return min

	// dp：获取总数
	//dp := make([]int, amount+1)
	//dp[0] = 1
	//for i := 1; i <= amount; i++ {
	//	for j := 0; j < len(coins); j++ {
	//		if i >= coins[j] {
	//			dp[i] += dp[i-coins[j]]
	//		}
	//	}
	//}
	//fmt.Println(dp)
	//return dp[amount]
}

func coinChangeRecursion(coins []int, amount int, change int, min *int) {
	fmt.Println(*min)
	if amount == 0 {
		*min = change
		return
	}
	if change >= *min-1 {
		return
	}
	for _, c := range coins {
		if amount >= c {
			coinChangeRecursion(coins, amount-c, change+1, min)
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
