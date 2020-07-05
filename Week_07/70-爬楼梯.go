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

import "fmt"

func main() {
	//n := 10
	//stairs := climbStairs(n)
	//fmt.Println(stairs)

	steps := []int{3, 7, 11}
	n := 99
	i := anySteps(steps, n)
	fmt.Println(i)

	repeat := anyStepsNoRepeat(steps, n)
	fmt.Println(repeat)
}

//leetcode submit region begin(Prohibit modification and deletion)
func climbStairs(n int) int {
	// 记忆化
	memo := make([]int, n+1)
	memo[0], memo[1] = 1, 2
	return climbStairsRecursion(n-1, memo)
}
func anySteps(steps []int, n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j < len(steps); j++ {
			if i >= steps[j] {
				dp[i] += dp[i-steps[j]]
			}
		}
	}
	return dp[n]
}
func anyStepsNoRepeat(steps []int, n int) int {
	dp, m := make([][]int, n+1), len(steps)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m)
	}
	for i := steps[0]; i <= n; i++ {
		for j := 0; j < m; j++ {
			if i > steps[j] {
				for k := 0; k < m; k++ {
					if j != k {
						dp[i][j] += dp[i-steps[j]][k]
					}
				}
			}
			if i == steps[j] {
				dp[i][j] ++
			}
		}
	}
	ans := 0
	for _, v := range dp[n] {
		ans += v
	}
	return ans
}
func climbStairsRecursion(n int, memo []int) int {
	if memo[n] == 0 {
		memo[n] = climbStairsRecursion(n-1, memo) + climbStairsRecursion(n-2, memo)
	}
	return memo[n]
}

//leetcode submit region end(Prohibit modification and deletion)
