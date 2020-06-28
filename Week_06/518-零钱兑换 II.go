//给定不同面额的硬币和一个总金额。写出函数来计算可以凑成总金额的硬币组合数。假设每一种面额的硬币有无限个。 
//
// 
//
// 
// 
//
// 示例 1: 
//
// 输入: amount = 5, coins = [1, 2, 5]
//输出: 4
//解释: 有四种方式可以凑成总金额:
//5=5
//5=2+2+1
//5=2+1+1+1
//5=1+1+1+1+1
// 
//
// 示例 2: 
//
// 输入: amount = 3, coins = [2]
//输出: 0
//解释: 只用面额2的硬币不能凑成总金额3。
// 
//
// 示例 3: 
//
// 输入: amount = 10, coins = [10] 
//输出: 1
// 
//
// 
//
// 注意: 
//
// 你可以假设： 
//
// 
// 0 <= amount (总金额) <= 5000 
// 1 <= coin (硬币面额) <= 5000 
// 硬币种类不超过 500 种 
// 结果符合 32 位符号整数 
// 
//
package main

import (
	"fmt"
)

func main() {
	amount, coins := 10, []int{3, 7, 11}
	i := change(amount, coins)
	fmt.Println(i)
}

//leetcode submit region begin(Prohibit modification and deletion)
func change(amount int, coins []int) int {
	// dp：
	//sort.Ints(coins)
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, c := range coins { // 巧妙之处：让零钱变得有“序”
		for i := 1; i <= amount; i++ {
			if i >= c {
				dp[i] += dp[i-c]
			}
		}
	}
	return dp[amount]

	// DFS：超时
	//return changeRecursion(amount, coins, amount+1)
}

func changeRecursion(amount int, coins []int, last int) int {
	if amount == 0 {
		return 1
	}
	ans := 0
	for _, c := range coins {
		if c > last || amount < c {
			continue
		}
		ans += changeRecursion(amount-c, coins, c)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
