//给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。 
//
// 
//
// 示例 1: 
//
// 输入: [2,3,-2,4]
//输出: 6
//解释: 子数组 [2,3] 有最大乘积 6。
// 
//
// 示例 2: 
//
// 输入: [-2,0,-1]
//输出: 0
//解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。 
// Related Topics 数组 动态规划
package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, -2, 4}
	product := maxProduct(nums)
	fmt.Println(product)
}

//leetcode submit region begin(Prohibit modification and deletion)
func maxProduct(nums []int) int {
	// dp：
	minAndMax := func(a, b int) (int, int) {
		if a > b {
			return a, b
		}
		return b, a
	}
	dpMax := make([]int, len(nums))
	dpMin := make([]int, len(nums))
	dpMax[0], dpMin[0] = nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		max, min := minAndMax(dpMax[i-1]*nums[i], dpMin[i-1]*nums[i])
		dpMax[i], _ = minAndMax(max, nums[i])
		_, dpMin[i] = minAndMax(min, nums[i])
	}
	max := dpMax[0]
	for i := 1; i < len(dpMax); i++ {
		if dpMax[i] > max {
			max = dpMax[i]
		}
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
