//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。 
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？ 
//
// 注意：给定 n 是一个正整数。 
//
// 示例 1： 
//
// 输入： 2
//输出： 2
//解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶 
//
// 示例 2： 
//
// 输入： 3
//输出： 3
//解释： 有三种方法可以爬到楼顶。
//1.  1 阶 + 1 阶 + 1 阶
//2.  1 阶 + 2 阶
//3.  2 阶 + 1 阶
// 
// Related Topics 动态规划
package main

import (
	"fmt"
)

func main() {
	n := 63
	stairs := climbStairs(n)
	fmt.Println(stairs)

	steps := []int{3, 7, 11}
	stepsClimbStairs := anyStepsClimbStairs(steps, n)
	fmt.Println("anySteps:", stepsClimbStairs)
}

//leetcode submit region begin(Prohibit modification and deletion)
func climbStairs(n int) int {
	// dp：
	one, two := 1, 1
	for i := 2; i <= n; i++ {
		one, two = two, one+two
	}
	return two
}
func anyStepsClimbStairs(steps []int, n int) int {
	// dp：任意步，相邻不相同
	sLen := len(steps)
	dp := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, sLen)
		for j := 0; j < sLen; j++ {
			if i > steps[j] {
				for k := 0; k < sLen; k++ {
					if k != j {
						dp[i][j] += dp[i-steps[j]][k]
					}
				}
			}
			if i == steps[j] {
				dp[i][j]++
			}
		}
	}
	ans := 0
	for _, d := range dp[n] {
		ans += d
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
