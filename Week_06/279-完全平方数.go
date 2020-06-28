//给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。 
//
// 示例 1: 
//
// 输入: n = 12
//输出: 3 
//解释: 12 = 4 + 4 + 4. 
//
// 示例 2: 
//
// 输入: n = 13
//输出: 2
//解释: 13 = 4 + 9. 
// Related Topics 广度优先搜索 数学 动态规划
package main

import "fmt"

func main() {
	n := 13
	squares := numSquares(n)
	fmt.Println(squares)
}

//leetcode submit region begin(Prohibit modification and deletion)
func numSquares(n int) int {
	// dp：先求开方
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	g := n
	for g*g > n {
		g = (g + n/g) >> 1
	}
	cache := make([]int, g)
	for i := 1; i <= g; i++ {
		cache[i-1] = i * i
	}
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = n + 1
	}
	for i := 1; i <= n; i++ {
		for j := 0; j < g; j++ {
			if i >= cache[j] {
				dp[i] = min(dp[i], dp[i-cache[j]]+1)
			}
		}
	}
	return dp[n]

	// 牛顿迭代法
	//g := n
	//for g*g > n {
	//	g = (g + n/g) >> 1
	//}
	//return g

	// 二分法
	//l, r := 1, n
	//ans := 0
	//for l <= r {
	//	mid := (l + r) >> 1
	//	val := mid * mid
	//	if val <= n {
	//		ans, l = mid, mid+1
	//	} else {
	//		r = mid - 1
	//	}
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
