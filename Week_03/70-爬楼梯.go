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
	"math"
)

func main() {
	//n := 11 // 144
	//n := 13 // 377
	n := 38 // 63245986
	stairs := climbStairs(n)
	fmt.Println(stairs)
}

/*
参考：https://www.cnblogs.com/benzhai/p/12410707.html
兔子繁殖问题：
	month 		1 2 3 4 5 ...
	count兔子 	1 2 3 5 8
	old 		1 1 2 3 5
	new 		0 1 1 2 3
	观察 old new 两行： 8 = 5 + 3
		5：理解为 上月老兔子 + 上月出生的新兔子 = 上月总和
		3：理解为 上月老兔子，而且 上月老兔子 = 上上月老兔子 + 上上月新兔子 = 上上月总和
	则有：F(n) = F(n-1) + F(n-2)
斐波那契公式：https://leetcode-cn.com/problems/climbing-stairs/solution/pa-lou-ti-by-leetcode/
*/
//leetcode submit region begin(Prohibit modification and deletion)
func climbStairs(n int) int {
	// 斐波那契公式
	sqrt5 := math.Sqrt(5)
	fibN := math.Pow((1+sqrt5)/2, float64(n+1)) - math.Pow((1-sqrt5)/2, float64(n+1))
	return int(math.Floor(fibN/sqrt5 + 0.5))

	// dp
	//two, one := 1, 1
	//for i := 1; i < n; i++ {
	//	two, one = one, two+one
	//}
	//return one

	// dp+数组
	//dp := make([]int, n)
	//dp[0], dp[1] = 1, 2
	//for i := 2; i < n; i++ {
	//	dp[i] = dp[i-1] + dp[i-2]
	//}
	//return dp[n-1]

	// 记忆化递归
	//memo := make([]int, n)
	//return climbStairsRecursion(memo, n)

	// 递归
	//if n < 2 {
	//	return 1
	//}
	//return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairsRecursion(memo []int, n int) int {
	if n < 2 {
		return 1
	}
	if memo[n-1] == 0 {
		memo[n-1] = climbStairsRecursion(memo, n-1) + climbStairsRecursion(memo, n-2)
	}
	return memo[n-1]
}

//leetcode submit region end(Prohibit modification and deletion)
