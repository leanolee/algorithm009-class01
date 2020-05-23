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
	n := 8
	stairs := climbStairs(n)
	fmt.Println(stairs)
}
/*
递归：O(2^n),O(1)
	1.终结条件：n < 3
	2.逻辑处理
	3.下一层：n-1，n-2
	4.清理当前层
缓存：O(n),O(n)
	思路：重复子结构的计算，大量冗余
dp：O(n),O(1)
	F(n) = F(n-1)+F(n-2)
*/
//leetcode submit region begin(Prohibit modification and deletion)
func climbStairs(n int) int {
	// 动态规划
	//stepFirst, stepSecond := 1, 1
	//for i := 2; i < n+1; i++ {
	//	stepFirst, stepSecond = stepSecond, stepFirst+stepSecond
	//}
	//return stepSecond

	// 缓存
	//memo := make([]int, n)
	//memo[0], memo[1] = 1, 2
	//return climbStairsRecursion(memo, n-1, n)

	// 递归
	if n < 3 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairsRecursion(memo []int, i int, n int) int {
	if memo[i] == 0 {
		memo[i] = climbStairsRecursion(memo, i-1, n) + climbStairsRecursion(memo, i-2, n)
	}
	return memo[i]
}

//leetcode submit region end(Prohibit modification and deletion)
