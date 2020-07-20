//有 n 个气球，编号为0 到 n-1，每个气球上都标有一个数字，这些数字存在数组 nums 中。 
//
// 现在要求你戳破所有的气球。如果你戳破气球 i ，就可以获得 nums[left] * nums[i] * nums[right] 个硬币。 这里的 lef
//t 和 right 代表和 i 相邻的两个气球的序号。注意当你戳破了气球 i 后，气球 left 和气球 right 就变成了相邻的气球。 
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
	//nums := []int{7, 9, 8, 0, 7, 1, 3, 5, 5, 2, 3}
	coins := maxCoins(nums)
	fmt.Println(coins)
}

/*
分治：
记忆化搜索：
	memo[l][r]：记录 l、r之间插入一个气球后的 coin
dp：自底向上
	cache := []int{1, nums[0], nums[1], ... nums[n-1], 1}
状态定义：
	dp[l][r]：
		l：左边界
		r：右边界
		dp[l][r]：l ~ r 之间的最大 coin
转移方程：
	dp[l][r] = max(dp[l][r], cache[l]*cache[i]*cache[r]+dp[l][i]+dp[i][r])
		l r 之间依次插入一个 cache[i] 值，求得 coin + dp[l][i] + dp[i][r])
		最后再从 l r 插入值对应的 coin 中求得最大值
	遍历方式一：从左上到右下，从底到顶
	for r := 1; r < n; r++ {
		for l := r - 2; l >= 0; l-- {
			for i := r - 1; i > l; i-- {
				dp[l][r] = max(balloons[l]*balloons[i]*balloons[r]+dp[l][i]+dp[i][r], dp[l][r])
			}
		}
	}
	遍历方式二：从右下到左上，从左到右
	for l := n - 1; l >= 0; l-- {
		for r := l + 2; r < n; r++ {
			for i := l + 1; i < r; i++ {
				dp[l][r] = max(dp[l][r], cache[l]*cache[i]*cache[r]+dp[l][i]+dp[i][r])
			}
		}
	}
*/
//leetcode submit region begin(Prohibit modification and deletion)
func maxCoins(nums []int) int {
	// dp：
	n := len(nums) + 2
	dp, cache := make([][]int, n), make([]int, n)
	cache[0], cache[n-1] = 1, 1
	copy(cache[1:], nums)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}
	for l := n - 1; l >= 0; l-- {
		for r := l + 2; r < n; r++ {
			for i := l + 1; i < r; i++ {
				dp[l][r] = max(dp[l][r], cache[l]*cache[i]*cache[r]+dp[l][i]+dp[i][r])
			}
			fmt.Println(l, r, dp)
		}
	}
	/*
		[[0 0 3 30 159 167]
		[0 0 0 15 135 159]
		[0 0 0 0 40 48]
		[0 0 0 0 0 40]
		[0 0 0 0 0 0]
		[0 0 0 0 0 0]]
	*/
	return dp[0][n-1]

	// 记忆化搜索：
	//n := len(nums)
	//memo, cache := make([][]int, n+2), make([]int, n+2)
	//cache[0], cache[n+1] = 1, 1
	//copy(cache[1:], nums)
	//for i := 0; i < len(memo); i++ {
	//	memo[i] = make([]int, n+2)
	//}
	//return coinsDFS(cache, memo, 0, n+1)

	// 分治：超时
	//coins, cache, max := make(map[int]bool), make([]int, len(nums)+2), 0
	//cache[0], cache[len(cache)-1] = 1, 1
	//copy(cache[1:], nums)
	//maxCoinsDFS(cache, &coins, 0)
	//for k, _ := range coins {
	//	if k > max {
	//		max = k
	//	}
	//}
	//return max
}

func coinsDFS(arr []int, memo [][]int, l int, r int) int {
	if l+1 >= r {
		return 0
	}
	if memo[l][r] == 0 {
		for i := l + 1; i < r; i++ {
			coin := coinsDFS(arr, memo, l, i) + arr[l]*arr[i]*arr[r] + coinsDFS(arr, memo, i, r)
			memo[l][r] = max(memo[l][r], coin)
		}
	}
	return memo[l][r]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxCoinsDFS(arr []int, ans *map[int]bool, coin int) {
	if len(arr) == 2 {
		(*ans)[coin] = true
		return
	}
	for i := 1; i < len(arr)-1; i++ {
		next := make([]int, len(arr)-1)
		copy(next, arr[:i])
		copy(next[i:], arr[i+1:])
		maxCoinsDFS(next, ans, coin+arr[i-1]*arr[i]*arr[i+1])
	}
}

//leetcode submit region end(Prohibit modification and deletion)
