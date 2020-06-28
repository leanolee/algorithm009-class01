//给定一个非负整数数组和一个整数 m，你需要将这个数组分成 m 个非空的连续子数组。设计一个算法使得这 m 个子数组各自和的最大值最小。 
//
// 注意: 
//数组长度 n 满足以下条件: 
//
// 
// 1 ≤ n ≤ 1000 
// 1 ≤ m ≤ min(50, n) 
// 
//
// 示例: 
//
// 
//输入:
//nums = [7,2,5,10,8]
//m = 2
//
//输出:
//18
//
//解释:
//一共有四种方法将nums分割为2个子数组。
//其中最好的方式是将其分为[7,2,5] 和 [10,8]，
//因为此时这两个子数组各自的和的最大值为18，在所有情况中最小。
// 
// Related Topics 二分查找 动态规划
package main

import (
	"fmt"
)

func main() {
	//nums := []int{7, 2, 5, 10, 8}
	nums := []int{2, 2147483647}
	m := 2
	array := splitArray(nums, m)
	fmt.Println(array)
}

/*
dp：
	首先我们把 f[i][j] 定义为将 nums[0..i] 分成 j 份时能得到的最小的分割子数组最大值。
	对于第 j 个子数组，它为数组中下标 k + 1 到 i 的这一段。因此，f[i][j] 可以从 max(f[k][j - 1], nums[k + 1] + ... + nums[i]) 这个公式中得到。遍历所有可能的 k，会得到 f[i][j] 的最小值。
	整个算法那的最终答案为 f[n][m]，其中 n 为数组大小。

*/
//leetcode submit region begin(Prohibit modification and deletion)
func splitArray(nums []int, m int) int {
	// 二分 + 贪心：
	minVal := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	l, r, n := 0, 0, len(nums)
	for i := 0; i < n; i++ {
		r += nums[i]
		if l < nums[i] {
			l = nums[i]
		}
	}
	//l = r / m	// 如果元素差太大，这样初始化就不ok
	fmt.Println(l,r)
	max := r
	for l <= r {
		mid, k, sum := (l+r)>>1, 1, 0
		for i := 0; i < n; i++ { // 保证所有 sum 都小于等于 mid
			if sum+nums[i] > mid {
				k++
				sum = nums[i]
			} else {
				sum += nums[i]
			}
		}
		fmt.Println(k, mid, max)
		if k <= m {
			max = minVal(max, mid)
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return max

	// dp：
	//minVal := func(a, b int) int {
	//	if a < b {
	//		return a
	//	}
	//	return b
	//}
	//maxVal := func(a, b int) int {
	//	if a > b {
	//		return a
	//	}
	//	return b
	//}
	//n := len(nums)
	//dp := make([][]int, n+1)
	//for i := 0; i <= n; i++ {
	//	dp[i] = make([]int, m+1)
	//	for j := 0; j <= m; j++ {
	//		dp[i][j] = math.MaxInt64
	//	}
	//}
	//sum := make([]int, n+1) // [0 7 9 14 24 32]
	//for i := 0; i < n; i++ {
	//	sum[i+1] = sum[i] + nums[i]
	//}
	//dp[0][0] = 0
	//for i := 1; i <= n; i++ {
	//	for j := 1; j <= m; j++ {
	//		for k := 0; k < i; k++ { // 上一个切分数组的位置
	//			dp[i][j] = minVal(maxVal(dp[k][j-1], sum[i]-sum[k]), dp[i][j]) // 最大最小
	//			//fmt.Println(dp, i, j)
	//		}
	//	}
	//}
	//return dp[n][m]

	// 暴力法：DFS，超时
	//max, n := math.MaxInt32, len(nums)
	//minVal := func(a, b int) int {
	//	if a < b {
	//		return a
	//	}
	//	return b
	//}
	//maxVal := func(a, b int) int {
	//	if a > b {
	//		return a
	//	}
	//	return b
	//}
	//var dfs func(sum int, currMax int, idx int, k int)
	//dfs = func(sum int, currMax int, idx int, k int) {
	//	if k == m && idx == n {
	//		max = minVal(max, currMax)
	//		return
	//	}
	//	if idx == n {
	//		return
	//	}
	//	if idx > 0 {
	//		dfs(sum+nums[idx], maxVal(currMax, sum+nums[idx]), idx+1, k)
	//	}
	//	if k < m {
	//		dfs(nums[idx], maxVal(currMax, nums[idx]), idx+1, k+1)
	//	}
	//}
	//dfs(0, 0, 0, 0)
	//return max
}

//leetcode submit region end(Prohibit modification and deletion)
