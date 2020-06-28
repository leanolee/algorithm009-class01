//有 n 个气球，编号为0 到 n-1，每个气球上都标有一个数字，这些数字存在数组 nums 中。 
//
// 现在要求你戳破所有的气球。每当你戳破一个气球 i 时，你可以获得 nums[left] * nums[i] * nums[right] 个硬币。 这里的 
//left 和 right 代表和 i 相邻的两个气球的序号。注意当你戳破了气球 i 后，气球 left 和气球 right 就变成了相邻的气球。 
//
// 求所能获得硬币的最大数量。 
//
// 说明: 
//
// 
// 你可以假设 nums[-1] = nums[n] = 1，但注意它们不是真实存在的所以并不能被戳破。 
// 0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100 
// 
//
// 示例: 
//
// 输入: [3,1,5,8]
//输出: 167 
//解释: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  --> []
//     coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167
// 
// Related Topics 分治算法 动态规划
package main

import "fmt"

func main() {
	nums := []int{3, 1, 5, 8}
	coins := maxCoins(nums)
	fmt.Println(coins)
}

//leetcode submit region begin(Prohibit modification and deletion)
func maxCoins(nums []int) int {
	// dp：自底向上
	n := len(nums) + 2
	balloons := make([]int, n)
	balloons[0], balloons[n-1] = 1, 1
	copy(balloons[1:], nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for r := 1; r < n; r++ {
		for l := r - 2; l >= 0; l-- {
			for i := r - 1; i > l; i-- {
				dp[l][r] = max(balloons[l]*balloons[i]*balloons[r]+dp[l][i]+dp[i][r], dp[l][r])
				//fmt.Println(dp)
			}
		}
	}
	return dp[0][n-1]

	// dp：自顶向下
	//n := len(nums) + 2
	//balloons := make([]int, n)
	//balloons[0], balloons[n-1] = 1, 1
	//copy(balloons[1:], nums)
	//dp := make([][]int, n)
	//for i := 0; i < n; i++ {
	//	dp[i] = make([]int, n)
	//}
	//max := func(a, b int) int {
	//	if a > b {
	//		return a
	//	}
	//	return b
	//}
	//var dfs func(l, r int) int
	//dfs = func(l, r int) int {
	//	if l+1 == r { // 不能添加气球了
	//		return 0
	//	}
	//	if dp[l][r] > 0 {
	//		return dp[l][r]
	//	}
	//	ans := 0
	//	for i := l + 1; i < r; i++ { // 向两者之间插入气球
	//		ans = max(balloons[l]*balloons[i]*balloons[r]+dfs(l, i)+dfs(i, r), ans)
	//	}
	//	dp[l][r] = ans
	//	return ans
	//}
	//return dfs(0, n-1)

	// 暴力法dfs：O(n!)
	return 0
}

//leetcode submit region end(Prohibit modification and deletion)
