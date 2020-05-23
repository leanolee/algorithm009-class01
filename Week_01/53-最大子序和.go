//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。 
//
// 示例: 
//
// 输入: [-2,1,-3,4,-1,2,1,-5,4],
//输出: 6
//解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
// 
//
// 进阶: 
//
// 如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。 
// Related Topics 数组 分治算法 动态规划
package main

import (
	"fmt"
)

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	//nums := []int{}
	array := maxSubArray(nums)
	fmt.Println(array)
}
/*
dp：O(n),O(1)
	F(n) = max(F(n-1)+nums[n], nums[n])
	如果 F(n-1) > 0; F(n) = F(n-1)+nums[n]
	否则 F(n) = nums[n]

	1.在原数组上操作，不推荐
		O(n),O(1)
	2.创建新数组记录 i 位置的最大值
		O(n),O(n)
	3.用2个变量记录 max、maxLast
		O(n),O(1)
*/
//leetcode submit region begin(Prohibit modification and deletion)
func maxSubArray(nums []int) int {
	// dp
	//n := len(nums)
	//dp := make([]int, n)
	//dp[0] = nums[0]
	//max := nums[0]
	//for i := 1; i < n; i++ {
	//	dp[i] = MyTools.Max(dp[i-1]+nums[i], nums[i])
	//	max = MyTools.Max(dp[i], max)
	//}
	//return max

	// dp：在原数组上操作，不推荐
	//n := len(nums)
	//if n == 0 {
	//	return 0
	//}
	//max := nums[0]
	//for i := 1; i < n; i++ {
	//	if nums[i-1] > 0 {
	//		nums[i] += nums[i-1]
	//	}
	//	if nums[i] > max {
	//		max = nums[i]
	//	}
	//}
	//return max

	// dp：用一个变量记录
	n := len(nums)
	if n == 0 {
		return 0
	}
	maxLast, max := nums[0], nums[0]
	for i := 1; i < n; i++ {
		if maxLast > 0 {
			maxLast += nums[i]
		} else {
			maxLast = nums[i]
		}
		if maxLast > max {
			max = maxLast
		}
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
