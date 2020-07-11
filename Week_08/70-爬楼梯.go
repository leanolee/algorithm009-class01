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
	n := 4
	stairs := climbStairs(n)
	fmt.Println(stairs)
}

//leetcode submit region begin(Prohibit modification and deletion)
func climbStairs(n int) [][]int {
	// dfs：记录路径，返回不同的方案集
	ans := make([][]int, 0)
	climbStairsArr(n, &ans, nil, 0)
	return ans
}

func climbStairsArr(n int, ans *[][]int, solve []int, i int) {
	if i > n {
		return
	}
	if i == n {
		*ans = append(*ans, solve)
		return
	}
	next := make([]int, len(solve))
	copy(next, solve)
	climbStairsArr(n, ans, append(next, 1), i+1)
	climbStairsArr(n, ans, append(next, 2), i+2)
}

//leetcode submit region end(Prohibit modification and deletion)
